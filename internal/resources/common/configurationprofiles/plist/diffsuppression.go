// common/configurationprofiles/plist/plistdiffsuppression.go
// contains the functions to process configuration profiles for diff suppression.
package plist

import (
	"html"
	"log"
	"strconv"
	"strings"

	"howett.net/plist"
)

// ProcessConfigurationProfileForDiffSuppression processes the plist data, removes specified fields, and returns the cleaned plist XML as a string.
func ProcessConfigurationProfileForDiffSuppression(plistData string, fieldsToRemove []string) (string, error) {
	log.Println("Starting ProcessConfigurationProfile")

	plistBytes := []byte(plistData)
	cleanedData, err := decodeAndCleanPlist(plistBytes, fieldsToRemove)
	if err != nil {
		log.Printf("Error decoding and cleaning plist data: %v\n", err)
		return "", err
	}

	// Normalize XML content in the plist
	normalizedData := normalizeXMLInPlist(cleanedData)

	sortedData := SortPlistKeys(normalizedData.(map[string]interface{}))

	log.Printf("Sorted and normalized plist data: %v\n", sortedData)

	// Encode the cleaned, normalized, and sorted data back to plist XML format
	encodedPlist, err := EncodePlist(sortedData)
	if err != nil {
		log.Printf("Error encoding cleaned data to plist: %v\n", err)
		return "", err
	}

	trimmedPlist := trimTrailingWhitespace(encodedPlist)

	return trimmedPlist, nil
}

// Function to decode a plist into a map and remove specified fields
func decodeAndCleanPlist(plistData []byte, fieldsToRemove []string) (map[string]interface{}, error) {
	var rawData map[string]interface{}
	_, err := plist.Unmarshal(plistData, &rawData)
	if err != nil {
		log.Printf("Error unmarshalling plist data: %v\n", err)
		return nil, err
	}

	log.Printf("Raw plist data: %v\n", rawData)
	RemoveFields(rawData, fieldsToRemove, "")
	log.Printf("Cleaned plist data: %v\n", rawData)

	return rawData, nil
}

// RemoveFields removes specified fields from a nested map
func RemoveFields(data map[string]interface{}, fieldsToRemove []string, path string) {
	// Create a set of fields to remove for quick lookup
	fieldsToRemoveSet := make(map[string]struct{}, len(fieldsToRemove))
	for _, field := range fieldsToRemove {
		fieldsToRemoveSet[field] = struct{}{}
	}

	recursivelyRemoveFields(data, fieldsToRemoveSet, path)
}

// recursivelyRemoveFields removes specified fields from a nested map
func recursivelyRemoveFields(data map[string]interface{}, fieldsToRemoveSet map[string]struct{}, path string) {
	// Iterate over the map and remove fields if they exist
	for field := range fieldsToRemoveSet {
		if _, exists := data[field]; exists {
			log.Printf("Removing field: %s from path: %s\n", field, path)
			delete(data, field)
		}
	}

	// Recursively process nested maps and arrays
	for key, value := range data {
		newPath := path + "/" + key
		switch v := value.(type) {
		case map[string]interface{}:
			log.Printf("Recursively removing fields in nested map at path: %s\n", newPath)
			recursivelyRemoveFields(v, fieldsToRemoveSet, newPath)
		case []interface{}:
			for i, item := range v {
				if nestedMap, ok := item.(map[string]interface{}); ok {
					log.Printf("Recursively removing fields in array at path: %s[%d]\n", newPath, i)
					recursivelyRemoveFields(nestedMap, fieldsToRemoveSet, newPath+strings.ReplaceAll(key, "/", "_")+strconv.Itoa(i))
				}
			}
			// Ensure empty arrays are preserved
			data[key] = v
		}
	}
}

// trimTrailingWhitespace removes trailing whitespace from each line of the plist
func trimTrailingWhitespace(plist string) string {
	lines := strings.Split(plist, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimRight(line, " \t")
	}
	return strings.Join(lines, "\n")
}

// normalizeXMLInPlist recursively normalizes XML content within a plist structure
// and unescapes HTML entities in string values to ensure consistent comparison.
func normalizeXMLInPlist(data interface{}) interface{} {
	switch v := data.(type) {
	case string:
		return html.UnescapeString(v)
	case map[string]interface{}:
		for key, value := range v {
			v[key] = normalizeXMLInPlist(value)
		}
	case []interface{}:
		for i, item := range v {
			v[i] = normalizeXMLInPlist(item)
		}
	}
	return data
}
