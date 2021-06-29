package aesolution

// GlobalAeopSkuProperty 
type GlobalAeopSkuProperty struct {
    // Customized attribute value name
    PropertyValueDefinitionName   string `json:"property_value_definition_name,omitempty" xml:"property_value_definition_name,omitempty"`
    // SKU attribute value id
    PropertyValueId   int64 `json:"property_value_id,omitempty" xml:"property_value_id,omitempty"`
    // Self-defined image url for this sku.
    SkuImage   string `json:"sku_image,omitempty" xml:"sku_image,omitempty"`
    // SKU attribute name id
    SkuPropertyId   int64 `json:"sku_property_id,omitempty" xml:"sku_property_id,omitempty"`
}
