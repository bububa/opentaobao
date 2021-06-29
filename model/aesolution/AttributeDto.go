package aesolution

// AttributeDTO 
type AttributeDTO struct {
    // merchant's attribute name
    AttributeName   string `json:"attribute_name,omitempty" xml:"attribute_name,omitempty"`
    // merchant's attribute value
    AttributeValue   string `json:"attribute_value,omitempty" xml:"attribute_value,omitempty"`
    // aliexpress attribute value id, which could be obtained from aliexpress.solution.sku.attribute.query
    AliexpressAttributeValueId   int64 `json:"aliexpress_attribute_value_id,omitempty" xml:"aliexpress_attribute_value_id,omitempty"`
    // aliexpress attribute name id, which could be obtained from aliexpress.solution.sku.attribute.query
    AliexpressAttributeNameId   int64 `json:"aliexpress_attribute_name_id,omitempty" xml:"aliexpress_attribute_name_id,omitempty"`
}
