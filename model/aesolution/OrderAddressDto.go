package aesolution

// OrderAddressDto 
type OrderAddressDto struct {
    // English country/region name
    CountryName   string `json:"country_name,omitempty" xml:"country_name,omitempty"`
    // mobile number
    MobileNo   string `json:"mobile_no,omitempty" xml:"mobile_no,omitempty"`
    // Recipient
    ContactPerson   string `json:"contact_person,omitempty" xml:"contact_person,omitempty"`
    // Phone country/region code
    PhoneCountry   string `json:"phone_country,omitempty" xml:"phone_country,omitempty"`
    // Phone area code
    PhoneArea   string `json:"phone_area,omitempty" xml:"phone_area,omitempty"`
    // province
    Province   string `json:"province,omitempty" xml:"province,omitempty"`
    // address 1
    Address   string `json:"address,omitempty" xml:"address,omitempty"`
    // phone number
    PhoneNumber   string `json:"phone_number,omitempty" xml:"phone_number,omitempty"`
    // Fax number
    FaxNumber   string `json:"fax_number,omitempty" xml:"fax_number,omitempty"`
    // Street detailed address
    DetailAddress   string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
    // city
    City   string `json:"city,omitempty" xml:"city,omitempty"`
    // country/region
    Country   string `json:"country,omitempty" xml:"country,omitempty"`
    // address 2
    Address2   string `json:"address2,omitempty" xml:"address2,omitempty"`
    // Fax country/region code
    FaxCountry   string `json:"fax_country,omitempty" xml:"fax_country,omitempty"`
    // Postal code
    Zip   string `json:"zip,omitempty" xml:"zip,omitempty"`
    // Fax area code
    FaxArea   string `json:"fax_area,omitempty" xml:"fax_area,omitempty"`
    // localized address, currently for buyer whose address is in Russia.
    LocalizedAddress   string `json:"localized_address,omitempty" xml:"localized_address,omitempty"`
}
