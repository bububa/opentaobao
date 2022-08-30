package aedropshiper

// AeItemSkuInfoDto 结构体
type AeItemSkuInfoDto struct {
	// SKU attribute object
	AeSkuPropertyDtos []AeSkuPropertyDto `json:"ae_sku_property_dtos,omitempty" xml:"ae_sku_property_dtos>ae_sku_property_dto,omitempty"`
	// Video information
	AeVideoDtos []AeVideoDto `json:"ae_video_dtos,omitempty" xml:"ae_video_dtos>ae_video_dto,omitempty"`
	// SKU ID
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// SKU price. Value range: 0.01-100000; Unit: USD. Such as: 200.07, which means: 200 US dollars 7 points. Need to be in the correct price range.
	SkuPrice string `json:"sku_price,omitempty" xml:"sku_price,omitempty"`
	// SKU merchant code. Format: single-byte alphanumeric characters, length 20, excluding spaces greater than and less than signs. If the user only fills in the retail price (productprice) and product code, a complete SKU record needs to be generated and submitted, otherwise the product code cannot be saved. The system will think that only the retail price has been submitted, but there is no SKU, resulting in unsaved product editing.
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// The currency unit of the product. U.S. Dollar: USD, Ruble: RUB
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// Commodity barcode
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// SKU discount price
	OfferSalePrice string `json:"offer_sale_price,omitempty" xml:"offer_sale_price,omitempty"`
	// SKU bulk discount price
	OfferBulkSalePrice string `json:"offer_bulk_sale_price,omitempty" xml:"offer_bulk_sale_price,omitempty"`
	// List of main images of the product
	ImageUrls string `json:"image_urls,omitempty" xml:"image_urls,omitempty"`
	// The actual saleable inventory attribute of SKU is ipmSkuStock. The reasonable value range of this attribute value is 0~999999. If the product has SKU, please make sure that at least one SKU is in stock, that is, the value of ipmSkuStock is 1~999999. The range of the inventory value of the entire product latitude is 1~999999. If the skuStock attribute is set at the same time, the system will give priority to the ipmSkuStock attribute; if the ipmSkuStock attribute is not set, the system will set the inventory according to the skuStock attribute, true means 999, false means 0.
	IpmSkuStock int64 `json:"ipm_sku_stock,omitempty" xml:"ipm_sku_stock,omitempty"`
	// Minimum number of batches
	SkuBulkOrder int64 `json:"sku_bulk_order,omitempty" xml:"sku_bulk_order,omitempty"`
	// SKU inventory
	SkuAvailableStock int64 `json:"sku_available_stock,omitempty" xml:"sku_available_stock,omitempty"`
	// SKU inventory, the data format is true if stock is available, false if no stock is available; at least one sku record is available.
	SkuStock bool `json:"sku_stock,omitempty" xml:"sku_stock,omitempty"`
}
