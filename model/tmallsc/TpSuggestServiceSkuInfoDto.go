package tmallsc

// TpSuggestServiceSkuInfoDto 结构体
type TpSuggestServiceSkuInfoDto struct {
	// 服务项
	ServiceSkuCode string `json:"service_sku_code,omitempty" xml:"service_sku_code,omitempty"`
	// 问题图片链接
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
}
