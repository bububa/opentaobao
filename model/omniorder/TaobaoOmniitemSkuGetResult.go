package omniorder

// TaobaoOmniitemSkuGetResult 
type TaobaoOmniitemSkuGetResult struct {
    // 错误码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 错误信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 返回查询到的sku列表
    Datas   []ItemLightPublishSkuDTO `json:"datas,omitempty" xml:"datas>item_light_publish_sku_dto,omitempty"`
}
