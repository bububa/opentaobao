package icbudropshipping

// Address 
type Address struct {
    // Specific address
    Address   string `json:"address,omitempty" xml:"address,omitempty"`
    // Secondary address
    AlternateAddress   string `json:"alternate_address,omitempty" xml:"alternate_address,omitempty"`
    // Name of the city where the order is completed
    City   string `json:"city,omitempty" xml:"city,omitempty"`
    // City abbreviation
    CityCode   string `json:"city_code,omitempty" xml:"city_code,omitempty"`
    // Name of contact person
    ContactPerson   string `json:"contact_person,omitempty" xml:"contact_person,omitempty"`
    // Country name
    Country   string `json:"country,omitempty" xml:"country,omitempty"`
    // Country code，ISO3166 standard and has two letters.
    CountryCode   string `json:"country_code,omitempty" xml:"country_code,omitempty"`
    // Fax
    Fax   *Phone `json:"fax,omitempty" xml:"fax,omitempty"`
    // Port name
    Port   string `json:"port,omitempty" xml:"port,omitempty"`
    // Port code
    PortCode   string `json:"port_code,omitempty" xml:"port_code,omitempty"`
    // Name of state/province
    Province   string `json:"province,omitempty" xml:"province,omitempty"`
    // Province/state abbreviation
    ProvinceCode   string `json:"province_code,omitempty" xml:"province_code,omitempty"`
    // telephone
    Telephone   *Phone `json:"telephone,omitempty" xml:"telephone,omitempty"`
    // zip
    Zip   string `json:"zip,omitempty" xml:"zip,omitempty"`
}
