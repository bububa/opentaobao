package aesolution

import (
	"sync"
)

// GlobalAeopAeProductSku 结构体
type GlobalAeopAeProductSku struct {
	// List of SKU attributes
	AeopSKUPropertyList []GlobalAeopSkuProperty `json:"aeop_s_k_u_property_list,omitempty" xml:"aeop_s_k_u_property_list>global_aeop_sku_property,omitempty"`
	// all of warehouse goods will return barcode
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// The Currency code. &#34;USD&#34; will be used as the default value if this information is not provided; Currency code is mandatory for Russian sellers(RUB) and Spanish sellers(EUR).
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// SKU ID. Can uniquely represent a SKU within a product range.
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// Sku merchant code from the seller&#39;s system. Format: alphanumeric, length 20, does not contain spaces greater than and less than sign. If you only fill in the product price and product code, you need to create a complete SKU record submission, otherwise the product code can not be saved. The system will consider that only the retail price is submitted, but no SKU, resulting in product editing is not saved.
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// Sku price. Value range: 0.01-100000; Such as: 200.07 means 200 US dollars 7 cents(if other currency_code is used, referring to the corresponding price in that currency, e.g., 200.07 Euros).
	SkuPrice string `json:"sku_price,omitempty" xml:"sku_price,omitempty"`
	// SKU discount price, also called sale price, value range: 0.01 - 100000.
	SkuDiscountPrice string `json:"sku_discount_price,omitempty" xml:"sku_discount_price,omitempty"`
	// EAN, or EAN13, stands for International Article Number (originally European Article Number). It is an extension of the UPC codes and you&#39;ll find them as barcodes on most everyday products. Sometimes the barcode is also called GTIN or GTIN13 (Global Trade Identifier)
	EanCode string `json:"ean_code,omitempty" xml:"ean_code,omitempty"`
	// Ranges from 1 to 999999 for one sku. The total stock of the entire product within multiple skus should also be in the range of 1 to 999999.
	IpmSkuStock int64 `json:"ipm_sku_stock,omitempty" xml:"ipm_sku_stock,omitempty"`
	// True means stock available for the sku, false means out of stock. The stock of at least one should be available.
	SkuStock bool `json:"sku_stock,omitempty" xml:"sku_stock,omitempty"`
}

var poolGlobalAeopAeProductSku = sync.Pool{
	New: func() any {
		return new(GlobalAeopAeProductSku)
	},
}

// GetGlobalAeopAeProductSku() 从对象池中获取GlobalAeopAeProductSku
func GetGlobalAeopAeProductSku() *GlobalAeopAeProductSku {
	return poolGlobalAeopAeProductSku.Get().(*GlobalAeopAeProductSku)
}

// ReleaseGlobalAeopAeProductSku 释放GlobalAeopAeProductSku
func ReleaseGlobalAeopAeProductSku(v *GlobalAeopAeProductSku) {
	v.AeopSKUPropertyList = v.AeopSKUPropertyList[:0]
	v.Barcode = ""
	v.CurrencyCode = ""
	v.Id = ""
	v.SkuCode = ""
	v.SkuPrice = ""
	v.SkuDiscountPrice = ""
	v.EanCode = ""
	v.IpmSkuStock = 0
	v.SkuStock = false
	poolGlobalAeopAeProductSku.Put(v)
}
