package util

// CountryDTO 
type CountryDTO struct {
    // 国家ID
    Id   string `json:"id,omitempty" xml:"id,omitempty"`
    // 国家名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 国家CODE
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
}
