package aesolution

// SkuAttributeDto 结构体
type SkuAttributeDto struct {
	// Deprecated, please use sku_attribute_name_id. To obtain the available sku attribute names under a specific category, please check API: aliexpress.solution.sku.attribute.query
	SkuAttributeName string `json:"sku_attribute_name,omitempty" xml:"sku_attribute_name,omitempty"`
	// Customized sku attribute value by sellers, do not include these 4 symbols { # : = , }.maximum 70 characters.
	SkuAttributeValue string `json:"sku_attribute_value,omitempty" xml:"sku_attribute_value,omitempty"`
	// Image that will represent the variation of the product. The url can point to a seller's server or to AliExpress photobank. In order to obtain more information about the photobank and how to upload images, please visit the following page: https://developers.aliexpress.com/en/doc.htm?docId=30186&docType=2
	SkuImage string `json:"sku_image,omitempty" xml:"sku_image,omitempty"`
	// Image that will represent the variation of the product. The url can point to a seller's server or to AliExpress photobank. In order to obtain more information about the photobank and how to upload images, please visit the following page: https://developers.aliexpress.com/en/doc.htm?docId=30186&docType=2
	SkuImageUrl string `json:"sku_image_url,omitempty" xml:"sku_image_url,omitempty"`
	// Please refer to https://developers.aliexpress.com/en/doc.htm?docId=29988&docType=2 to obtain the sku_attribute_value_id under specific category
	SkuAttributeValueId int64 `json:"sku_attribute_value_id,omitempty" xml:"sku_attribute_value_id,omitempty"`
	// Please refer to https://developers.aliexpress.com/en/doc.htm?docId=29988&docType=2 to obtain the sku_attribute_name_id under specific category
	SkuAttributeNameId int64 `json:"sku_attribute_name_id,omitempty" xml:"sku_attribute_name_id,omitempty"`
}
