---
page_title: "jamfpro_computer_inventory"
description: |-
  
---

# jamfpro_computer_inventory (Data Source)


<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `applications` (List of Object) (see [below for nested schema](#nestedatt--applications))
- `attachments` (List of Object) (see [below for nested schema](#nestedatt--attachments))
- `certificates` (List of Object) (see [below for nested schema](#nestedatt--certificates))
- `configuration_profiles` (List of Object) (see [below for nested schema](#nestedatt--configuration_profiles))
- `disk_encryption` (List of Object) (see [below for nested schema](#nestedatt--disk_encryption))
- `extension_attributes` (List of Object) (see [below for nested schema](#nestedatt--extension_attributes))
- `fonts` (List of Object) (see [below for nested schema](#nestedatt--fonts))
- `general` (List of Object) (see [below for nested schema](#nestedatt--general))
- `group_memberships` (List of Object) (see [below for nested schema](#nestedatt--group_memberships))
- `hardware` (List of Object) (see [below for nested schema](#nestedatt--hardware))
- `ibeacons` (List of Object) (see [below for nested schema](#nestedatt--ibeacons))
- `id` (String) The ID of this resource.
- `licensed_software` (List of Object) (see [below for nested schema](#nestedatt--licensed_software))
- `local_user_accounts` (List of Object) (see [below for nested schema](#nestedatt--local_user_accounts))
- `operating_system` (List of Object) (see [below for nested schema](#nestedatt--operating_system))
- `package_receipts` (List of Object) (see [below for nested schema](#nestedatt--package_receipts))
- `plugins` (List of Object) (see [below for nested schema](#nestedatt--plugins))
- `printers` (List of Object) (see [below for nested schema](#nestedatt--printers))
- `purchasing` (List of Object) (see [below for nested schema](#nestedatt--purchasing))
- `security` (List of Object) (see [below for nested schema](#nestedatt--security))
- `services` (List of Object) (see [below for nested schema](#nestedatt--services))
- `software_updates` (List of Object) (see [below for nested schema](#nestedatt--software_updates))
- `storage` (List of Object) (see [below for nested schema](#nestedatt--storage))
- `udid` (String)
- `user_and_location` (List of Object) (see [below for nested schema](#nestedatt--user_and_location))

<a id="nestedatt--applications"></a>
### Nested Schema for `applications`

Read-Only:

- `bundle_id` (String)
- `external_version_id` (String)
- `mac_app_store` (Boolean)
- `name` (String)
- `path` (String)
- `size_megabytes` (Number)
- `update_available` (Boolean)
- `version` (String)


<a id="nestedatt--attachments"></a>
### Nested Schema for `attachments`

Read-Only:

- `file_type` (String)
- `id` (String)
- `name` (String)
- `size_bytes` (Number)


<a id="nestedatt--certificates"></a>
### Nested Schema for `certificates`

Read-Only:

- `certificate_status` (String)
- `common_name` (String)
- `expiration_date` (String)
- `identity` (Boolean)
- `issued_date` (String)
- `lifecycle_status` (String)
- `serial_number` (String)
- `sha1_fingerprint` (String)
- `subject_name` (String)
- `username` (String)


<a id="nestedatt--configuration_profiles"></a>
### Nested Schema for `configuration_profiles`

Read-Only:

- `display_name` (String)
- `id` (String)
- `last_installed` (String)
- `profile_identifier` (String)
- `removable` (Boolean)
- `username` (String)


<a id="nestedatt--disk_encryption"></a>
### Nested Schema for `disk_encryption`

Read-Only:

- `boot_partition_encryption_details` (List of Object) (see [below for nested schema](#nestedobjatt--disk_encryption--boot_partition_encryption_details))
- `disk_encryption_configuration_name` (String)
- `file_vault2_eligibility_message` (String)
- `file_vault2_enabled_user_names` (List of String)
- `individual_recovery_key_validity_status` (String)
- `institutional_recovery_key_present` (Boolean)

<a id="nestedobjatt--disk_encryption--boot_partition_encryption_details"></a>
### Nested Schema for `disk_encryption.boot_partition_encryption_details`

Read-Only:

- `partition_file_vault2_percent` (Number)
- `partition_file_vault2_state` (String)
- `partition_name` (String)



<a id="nestedatt--extension_attributes"></a>
### Nested Schema for `extension_attributes`

Read-Only:

- `data_type` (String)
- `definition_id` (String)
- `description` (String)
- `enabled` (Boolean)
- `input_type` (String)
- `multi_value` (Boolean)
- `name` (String)
- `options` (List of String)
- `values` (List of String)


<a id="nestedatt--fonts"></a>
### Nested Schema for `fonts`

Read-Only:

- `name` (String)
- `path` (String)
- `version` (String)


<a id="nestedatt--general"></a>
### Nested Schema for `general`

Read-Only:

- `asset_tag` (String)
- `barcode1` (String)
- `barcode2` (String)
- `declarative_device_management_enabled` (Boolean)
- `distribution_point` (String)
- `enrolled_via_automated_device_enrollment` (Boolean)
- `enrollment_method` (List of Object) (see [below for nested schema](#nestedobjatt--general--enrollment_method))
- `extension_attributes` (List of Object) (see [below for nested schema](#nestedobjatt--general--extension_attributes))
- `initial_entry_date` (String)
- `itunes_store_account_active` (Boolean)
- `jamf_binary_version` (String)
- `last_cloud_backup_date` (String)
- `last_contact_time` (String)
- `last_enrolled_date` (String)
- `last_ip_address` (String)
- `last_reported_ip` (String)
- `management_id` (String)
- `mdm_capable` (List of Object) (see [below for nested schema](#nestedobjatt--general--mdm_capable))
- `mdm_profile_expiration` (String)
- `name` (String)
- `platform` (String)
- `remote_management` (List of Object) (see [below for nested schema](#nestedobjatt--general--remote_management))
- `report_date` (String)
- `site_id` (List of Object) (see [below for nested schema](#nestedobjatt--general--site_id))
- `supervised` (Boolean)
- `user_approved_mdm` (Boolean)

<a id="nestedobjatt--general--enrollment_method"></a>
### Nested Schema for `general.enrollment_method`

Read-Only:

- `id` (String)
- `object_name` (String)
- `object_type` (String)


<a id="nestedobjatt--general--extension_attributes"></a>
### Nested Schema for `general.extension_attributes`

Read-Only:

- `data_type` (String)
- `definition_id` (String)
- `description` (String)
- `enabled` (Boolean)
- `input_type` (String)
- `multi_value` (Boolean)
- `name` (String)
- `options` (List of String)
- `values` (List of String)


<a id="nestedobjatt--general--mdm_capable"></a>
### Nested Schema for `general.mdm_capable`

Read-Only:

- `capable` (Boolean)
- `capable_users` (List of String)


<a id="nestedobjatt--general--remote_management"></a>
### Nested Schema for `general.remote_management`

Read-Only:

- `managed` (Boolean)
- `management_username` (String)


<a id="nestedobjatt--general--site_id"></a>
### Nested Schema for `general.site_id`

Read-Only:

- `id` (String)
- `name` (String)



<a id="nestedatt--group_memberships"></a>
### Nested Schema for `group_memberships`

Read-Only:

- `group_id` (String)
- `group_name` (String)
- `smart_group` (Boolean)


<a id="nestedatt--hardware"></a>
### Nested Schema for `hardware`

Read-Only:

- `alt_mac_address` (String)
- `alt_network_adapter_type` (String)
- `apple_silicon` (Boolean)
- `battery_capacity_percent` (Number)
- `ble_capable` (Boolean)
- `boot_rom` (String)
- `bus_speed_mhz` (Number)
- `cache_size_kilobytes` (Number)
- `core_count` (Number)
- `extension_attributes` (List of Object) (see [below for nested schema](#nestedobjatt--hardware--extension_attributes))
- `mac_address` (String)
- `make` (String)
- `model` (String)
- `model_identifier` (String)
- `network_adapter_type` (String)
- `nic_speed` (String)
- `open_ram_slots` (Number)
- `optical_drive` (String)
- `processor_architecture` (String)
- `processor_count` (Number)
- `processor_speed_mhz` (Number)
- `processor_type` (String)
- `serial_number` (String)
- `smc_version` (String)
- `supports_ios_app_installs` (Boolean)
- `total_ram_megabytes` (Number)

<a id="nestedobjatt--hardware--extension_attributes"></a>
### Nested Schema for `hardware.extension_attributes`

Read-Only:

- `data_type` (String)
- `definition_id` (String)
- `description` (String)
- `enabled` (Boolean)
- `input_type` (String)
- `multi_value` (Boolean)
- `name` (String)
- `options` (List of String)
- `values` (List of String)



<a id="nestedatt--ibeacons"></a>
### Nested Schema for `ibeacons`

Read-Only:

- `name` (String)


<a id="nestedatt--licensed_software"></a>
### Nested Schema for `licensed_software`

Read-Only:

- `id` (String)
- `name` (String)


<a id="nestedatt--local_user_accounts"></a>
### Nested Schema for `local_user_accounts`

Read-Only:

- `admin` (Boolean)
- `azure_active_directory_id` (String)
- `computer_azure_active_directory_id` (String)
- `file_vault2_enabled` (Boolean)
- `full_name` (String)
- `home_directory` (String)
- `home_directory_size_mb` (Number)
- `password_history_depth` (Number)
- `password_max_age` (Number)
- `password_min_complex_characters` (Number)
- `password_min_length` (Number)
- `password_require_alphanumeric` (Boolean)
- `uid` (String)
- `user_account_type` (String)
- `user_azure_active_directory_id` (String)
- `user_guid` (String)
- `username` (String)


<a id="nestedatt--operating_system"></a>
### Nested Schema for `operating_system`

Read-Only:

- `active_directory_status` (String)
- `build` (String)
- `extension_attributes` (List of Object) (see [below for nested schema](#nestedobjatt--operating_system--extension_attributes))
- `filevault2_status` (String)
- `name` (String)
- `rapid_security_response` (String)
- `software_update_device_id` (String)
- `supplemental_build_version` (String)
- `version` (String)

<a id="nestedobjatt--operating_system--extension_attributes"></a>
### Nested Schema for `operating_system.extension_attributes`

Read-Only:

- `data_type` (String)
- `definition_id` (String)
- `description` (String)
- `enabled` (Boolean)
- `input_type` (String)
- `multi_value` (Boolean)
- `name` (String)
- `options` (List of String)
- `values` (List of String)



<a id="nestedatt--package_receipts"></a>
### Nested Schema for `package_receipts`

Read-Only:

- `cached` (List of String)
- `installed_by_installer_swu` (List of String)
- `installed_by_jamf_pro` (List of String)


<a id="nestedatt--plugins"></a>
### Nested Schema for `plugins`

Read-Only:

- `name` (String)
- `path` (String)
- `version` (String)


<a id="nestedatt--printers"></a>
### Nested Schema for `printers`

Read-Only:

- `location` (String)
- `name` (String)
- `type` (String)
- `uri` (String)


<a id="nestedatt--purchasing"></a>
### Nested Schema for `purchasing`

Read-Only:

- `apple_care_id` (String)
- `extension_attributes` (List of Object) (see [below for nested schema](#nestedobjatt--purchasing--extension_attributes))
- `lease_date` (String)
- `leased` (Boolean)
- `life_expectancy` (Number)
- `po_date` (String)
- `po_number` (String)
- `purchase_price` (String)
- `purchased` (Boolean)
- `purchasing_account` (String)
- `purchasing_contact` (String)
- `vendor` (String)
- `warranty_date` (String)

<a id="nestedobjatt--purchasing--extension_attributes"></a>
### Nested Schema for `purchasing.extension_attributes`

Read-Only:

- `data_type` (String)
- `definition_id` (String)
- `description` (String)
- `enabled` (Boolean)
- `input_type` (String)
- `multi_value` (Boolean)
- `name` (String)
- `options` (List of String)
- `values` (List of String)



<a id="nestedatt--security"></a>
### Nested Schema for `security`

Read-Only:

- `activation_lock_enabled` (Boolean)
- `auto_login_disabled` (Boolean)
- `bootstrap_token_allowed` (Boolean)
- `external_boot_level` (String)
- `firewall_enabled` (Boolean)
- `gatekeeper_status` (String)
- `recovery_lock_enabled` (Boolean)
- `remote_desktop_enabled` (Boolean)
- `secure_boot_level` (String)
- `sip_status` (String)
- `xprotect_version` (String)


<a id="nestedatt--services"></a>
### Nested Schema for `services`

Read-Only:

- `name` (String)


<a id="nestedatt--software_updates"></a>
### Nested Schema for `software_updates`

Read-Only:

- `name` (String)
- `package_name` (String)
- `version` (String)


<a id="nestedatt--storage"></a>
### Nested Schema for `storage`

Read-Only:

- `boot_drive_available_space_megabytes` (Number)
- `disks` (List of Object) (see [below for nested schema](#nestedobjatt--storage--disks))

<a id="nestedobjatt--storage--disks"></a>
### Nested Schema for `storage.disks`

Read-Only:

- `device` (String)
- `id` (String)
- `model` (String)
- `partitions` (List of Object) (see [below for nested schema](#nestedobjatt--storage--disks--partitions))
- `revision` (String)
- `serial_number` (String)
- `size_megabytes` (Number)
- `smart_status` (String)
- `type` (String)

<a id="nestedobjatt--storage--disks--partitions"></a>
### Nested Schema for `storage.disks.partitions`

Read-Only:

- `available_megabytes` (Number)
- `file_vault2_progress_percent` (Number)
- `file_vault2_state` (String)
- `lvm_managed` (Boolean)
- `name` (String)
- `partition_type` (String)
- `percent_used` (Number)
- `size_megabytes` (Number)




<a id="nestedatt--user_and_location"></a>
### Nested Schema for `user_and_location`

Read-Only:

- `building_id` (String)
- `department_id` (String)
- `email` (String)
- `extension_attributes` (List of Object) (see [below for nested schema](#nestedobjatt--user_and_location--extension_attributes))
- `phone` (String)
- `position` (String)
- `realname` (String)
- `room` (String)
- `username` (String)

<a id="nestedobjatt--user_and_location--extension_attributes"></a>
### Nested Schema for `user_and_location.extension_attributes`

Read-Only:

- `data_type` (String)
- `definition_id` (String)
- `description` (String)
- `enabled` (Boolean)
- `input_type` (String)
- `multi_value` (Boolean)
- `name` (String)
- `options` (List of String)
- `values` (List of String)