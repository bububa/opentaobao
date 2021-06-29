package aesolution

// SupportedSkuAttributeDto 
type SupportedSkuAttributeDto struct {

    // aliexpress sku name, the same field when indicating the sku_name for posting product
    
    AliexpressSkuName   string `json:"aliexpress_sku_name,omitempty" xml:"aliexpress_sku_name,omitempty"`
    

    // Indicates whether this sku attribute is mandatory under this category
    
    Required   bool `json:"required,omitempty" xml:"required,omitempty"`
    

    // aliexpress sku value list
    
    AliexpressSkuValueList   []SkuValueSimplifiedInfoDto `json:"aliexpress_sku_value_list,omitempty" xml:"aliexpress_sku_value_list,omitempty"`
    

    // whether the corresponding aliexpress_sku_name support customized name by merchants
    
    SupportCustomizedName   bool `json:"support_customized_name,omitempty" xml:"support_customized_name,omitempty"`
    

    // whether the corresponding aliexpress_sku_name support customized picture
    
    SupportCustomizedPicture   bool `json:"support_customized_picture,omitempty" xml:"support_customized_picture,omitempty"`
    

}
