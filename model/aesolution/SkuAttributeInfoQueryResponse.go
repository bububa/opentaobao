package aesolution

// SkuAttributeInfoQueryResponse 
type SkuAttributeInfoQueryResponse struct {
    // supported sku attribute lis
    SupportingSkuAttributeList   []SupportedSkuAttributeDto `json:"supporting_sku_attribute_list,omitempty" xml:"supporting_sku_attribute_list>supported_sku_attribute_dto,omitempty"`
    // common attributes under a specific category
    SupportingCommonAttributeList   []SupportedCommonAttributeDto `json:"supporting_common_attribute_list,omitempty" xml:"supporting_common_attribute_list>supported_common_attribute_dto,omitempty"`
}
