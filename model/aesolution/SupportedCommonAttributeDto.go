package aesolution

// SupportedCommonAttributeDto 
type SupportedCommonAttributeDto struct {
    // aliexpress common attribute name id
    AliexpressCommonAttributeNameId   int64 `json:"aliexpress_common_attribute_name_id,omitempty" xml:"aliexpress_common_attribute_name_id,omitempty"`
    // aliexpress common attribute name
    AliexpressCommonAttributeName   string `json:"aliexpress_common_attribute_name,omitempty" xml:"aliexpress_common_attribute_name,omitempty"`
    // whether the common attribute is required under this category
    Required   bool `json:"required,omitempty" xml:"required,omitempty"`
    // aliexpress common attribute value list
    AliexpressCommonAttributeValueList   []CommonAttributeValueInfoDto `json:"aliexpress_common_attribute_value_list,omitempty" xml:"aliexpress_common_attribute_value_list>common_attribute_value_info_dto,omitempty"`
}
