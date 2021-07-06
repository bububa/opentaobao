package util

// AssetQrCodeDto 结构体
type AssetQrCodeDto struct {
	// 生产二维码信息字符串
	QrCodeStringList []string `json:"qr_code_string_list,omitempty" xml:"qr_code_string_list>string,omitempty"`
	// 资产类型
	AssetType string `json:"asset_type,omitempty" xml:"asset_type,omitempty"`
	// 实物来源
	EntitySource string `json:"entity_source,omitempty" xml:"entity_source,omitempty"`
	// 请求生产条码数量
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 配件厂商code(请求参数)
	SpareBrandCode string `json:"spare_brand_code,omitempty" xml:"spare_brand_code,omitempty"`
	// 整机厂商code(请求参数)
	DeviceBrandCode string `json:"device_brand_code,omitempty" xml:"device_brand_code,omitempty"`
	// 实物SN(请求参数)
	Sn string `json:"sn,omitempty" xml:"sn,omitempty"`
	// 阿里侧部件型号(请求参数)
	Mpn string `json:"mpn,omitempty" xml:"mpn,omitempty"`
	// 配件类型code(请求参数)
	SpareCategoryCode string `json:"spare_category_code,omitempty" xml:"spare_category_code,omitempty"`
	// 厂商代码(请求参数)
	Manufacture string `json:"manufacture,omitempty" xml:"manufacture,omitempty"`
	// 单批次请求唯一标识
	Nonce string `json:"nonce,omitempty" xml:"nonce,omitempty"`
}
