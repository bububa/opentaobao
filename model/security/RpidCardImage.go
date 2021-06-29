package security

// RpidCardImage 
type RpidCardImage struct {
    // backImageUrl
    BackImageUrl   string `json:"back_image_url,omitempty" xml:"back_image_url,omitempty"`
    // cardType
    CardType   string `json:"card_type,omitempty" xml:"card_type,omitempty"`
    // frontImageUrl
    FrontImageUrl   string `json:"front_image_url,omitempty" xml:"front_image_url,omitempty"`
}
