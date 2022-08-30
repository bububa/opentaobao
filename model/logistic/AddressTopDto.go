package logistic

// AddressTopDto 结构体
type AddressTopDto struct {
	// first name of sender
	FirstName string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// last name of sender
	LastName string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	// sender's country
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// zip code of ship from place
	ZipCode string `json:"zip_code,omitempty" xml:"zip_code,omitempty"`
	// sender's district and street
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// federal tax id(cnpj)
	FederalTaxId string `json:"federal_tax_id,omitempty" xml:"federal_tax_id,omitempty"`
	// sender's city
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// sender's street number
	AddressNumber string `json:"address_number,omitempty" xml:"address_number,omitempty"`
	// cell phone of sender
	Cellphone string `json:"cellphone,omitempty" xml:"cellphone,omitempty"`
	// sender's State
	State string `json:"state,omitempty" xml:"state,omitempty"`
	// email of sender
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// shipping additional
	Additional string `json:"additional,omitempty" xml:"additional,omitempty"`
}
