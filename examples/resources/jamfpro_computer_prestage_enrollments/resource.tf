// Minimum Configuration
resource "jamfpro_computer_prestage_enrollment" "example" {
  display_name                            = "jamfpro-sdk-example-computerPrestageMinimum-config"
  mandatory                               = true
  mdm_removable                           = true
  support_phone_number                    = "111-222-3333"
  support_email_address                   = "email@company.com"
  department                              = "department name"
  default_prestage                        = false
  enrollment_site_id                      = "-1"
  keep_existing_site_membership           = false
  keep_existing_location_information      = false
  require_authentication                  = false
  authentication_prompt                   = "hello welcome to your enterprise managed macOS device"
  prevent_activation_lock                 = false
  enable_device_based_activation_lock     = false
  device_enrollment_program_instance_id   = "1"
  skip_setup_items {
    biometric           = false
    terms_of_address    = false
    file_vault          = false
    icloud_diagnostics  = false
    diagnostics         = false
    accessibility       = false
    apple_id            = false
    screen_time         = false
    siri                = false
    display_tone        = false
    restore             = false
    appearance          = false
    privacy             = false
    payment             = false
    registration        = false
    tos                 = false
    icloud_storage      = false
    location            = false
  }
  location_information {
    username       = ""
    realname       = ""
    phone          = ""
    email          = ""
    room           = ""
    position       = ""
    department_id  = "-1"
    building_id    = "-1"
  }
  purchasing_information {
    leased              = false
    purchased           = true
    apple_care_id       = ""
    po_number           = ""
    vendor              = ""
    purchase_price      = ""
    life_expectancy     = 0
    purchasing_account  = ""
    purchasing_contact  = ""
    lease_date          = "1970-01-01"
    po_date             = "1970-01-01"
    warranty_date       = "1970-01-01"
  }
  anchor_certificates                     = []
  enrollment_customization_id             = "0"
  language                                = ""
  region                                  = ""
  auto_advance_setup                      = false
  install_profiles_during_setup           = true
  prestage_installed_profile_ids          = []
  custom_package_ids                      = []
  custom_package_distribution_point_id    = "-1"
  enable_recovery_lock                    = false
  recovery_lock_password_type             = "MANUAL"
  recovery_lock_password                  = ""
  rotate_recovery_lock_password           = false
  site_id                                 = "-1"
  account_settings {
    payload_configured                            = true
    local_admin_account_enabled                   = false
    admin_username                                = ""
    admin_password                                = ""
    hidden_admin_account                          = false
    local_user_managed                            = false
    user_account_type                             = "ADMINISTRATOR"
    prefill_primary_account_info_feature_enabled  = false
    prefill_type                                  = "UNKNOWN"
    prefill_account_full_name                     = ""
    prefill_account_user_name                     = ""
    prevent_prefill_info_from_modification        = false
  }
}

// Configure the Jamf Pro Computer Prestage Enrollment

resource "jamfpro_computer_prestage_enrollment" "jamf_pro_computer_prestage_enrollment_001" {
  display_name                          = "tf-computer-prestage-enrollment-001"
  mandatory                             = true
  mdm_removable                         = true
  support_phone_number                  = "111-222-3333"
  support_email_address                 = "email@company.com"
  department                            = "department name"
  default_prestage                      = false
  enrollment_site_id                    = "-1"
  keep_existing_site_membership         = false
  keep_existing_location_information    = false
  require_authentication                = false
  authentication_prompt                 = "hello welcome to your enterprise managed macOS device"
  prevent_activation_lock               = false
  enable_device_based_activation_lock   = false
  device_enrollment_program_instance_id = "1"
  skip_setup_items {
    biometric          = false
    terms_of_address   = false
    file_vault         = false
    icloud_diagnostics = false
    diagnostics        = false
    accessibility      = false
    apple_id           = false
    screen_time        = false
    siri               = false
    display_tone       = false
    restore            = false
    appearance         = false
    privacy            = false
    payment            = false
    registration       = false
    tos                = false
    icloud_storage     = false
    location           = false
  }
  location_information {
    username      = ""
    realname      = ""
    phone         = ""
    email         = ""
    room          = ""
    position      = ""
    department_id = "-1"
    building_id   = "-1"
  }
  purchasing_information {
    leased             = false
    purchased          = true
    apple_care_id      = ""
    po_number          = ""
    vendor             = ""
    purchase_price     = ""
    life_expectancy    = 0
    purchasing_account = ""
    purchasing_contact = ""
    lease_date         = "1970-01-01"
    po_date            = "1970-01-01"
    warranty_date      = "1970-01-01"
  }
  anchor_certificates                  = []
  enrollment_customization_id          = "0"
  auto_advance_setup                   = false
  language                             = "" // en
  region                               = "" // GB
  install_profiles_during_setup        = true
  prestage_installed_profile_ids       = reverse(sort([3114, 3800])) // requires decending order
  custom_package_ids                   = sort([1, 2]) // requires ascending order
  custom_package_distribution_point_id = "-2" // "-1" - not used / "-2" - Cloud Distribution Point (Jamf Cloud) / "any other number" - Distribution Point ID
  enable_recovery_lock                 = false
  recovery_lock_password_type          = "MANUAL"
  recovery_lock_password               = ""
  rotate_recovery_lock_password        = false
  site_id                              = "-1"
  account_settings {
    payload_configured                           = true
    local_admin_account_enabled                  = true
    admin_username                               = "thing"
    admin_password                               = "thing"
    hidden_admin_account                         = true
    local_user_managed                           = true
    user_account_type                            = "ADMINISTRATOR" // "STANDARD" / "ADMINISTRATOR" / "SKIP"
    prefill_primary_account_info_feature_enabled = true
    prefill_type                                 = "CUSTOM" // "UNKNOWN" / "CUSTOM" / "DEVICE_OWNER"
    prefill_account_full_name                    = "firstname.lastname"
    prefill_account_user_name                    = "firstname.lastname"
    prevent_prefill_info_from_modification       = false
  }
}