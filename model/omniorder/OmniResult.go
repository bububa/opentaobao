package omniorder

// OmniResult 结构体
type OmniResult struct {
	// 返回查询到的sku列表
	Datas []ItemLightPublishSkuDto `json:"datas,omitempty" xml:"datas>item_light_publish_sku_dto,omitempty"`
	// 错误码，0表示成功
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// data
	Data *ItemDeleteResult `json:"data,omitempty" xml:"data,omitempty"`
}
