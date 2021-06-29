package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
全渠道门店商品轻发布 API请求
taobao.omniitem.item.publish

全渠道门店商品轻发布
*/
type TaobaoOmniitemItemPublishRequest struct {
    model.Params
    // 发布商品信息
    lightPublishInfo   *ItemLightPublishDto
    // 在全域商品或是门店商品中校验码是否重复，可选值对应为ALL或者STORE
    operateType   string
}

// 初始化TaobaoOmniitemItemPublishRequest对象
func NewTaobaoOmniitemItemPublishRequest() *TaobaoOmniitemItemPublishRequest{
    return &TaobaoOmniitemItemPublishRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniitemItemPublishRequest) GetApiMethodName() string {
    return "taobao.omniitem.item.publish"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniitemItemPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LightPublishInfo Setter
// 发布商品信息
func (r *TaobaoOmniitemItemPublishRequest) SetLightPublishInfo(lightPublishInfo *ItemLightPublishDto) error {
    r.lightPublishInfo = lightPublishInfo
    r.Set("light_publish_info", lightPublishInfo)
    return nil
}

// LightPublishInfo Getter
func (r TaobaoOmniitemItemPublishRequest) GetLightPublishInfo() *ItemLightPublishDto {
    return r.lightPublishInfo
}
// OperateType Setter
// 在全域商品或是门店商品中校验码是否重复，可选值对应为ALL或者STORE
func (r *TaobaoOmniitemItemPublishRequest) SetOperateType(operateType string) error {
    r.operateType = operateType
    r.Set("operate_type", operateType)
    return nil
}

// OperateType Getter
func (r TaobaoOmniitemItemPublishRequest) GetOperateType() string {
    return r.operateType
}
