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
type TaobaoOmniitemItemPublishAPIRequest struct {
    model.Params
    // 发布商品信息
    _lightPublishInfo   *ItemLightPublishDTO
    // 在全域商品或是门店商品中校验码是否重复，可选值对应为ALL或者STORE
    _operateType   string
}

// 初始化TaobaoOmniitemItemPublishAPIRequest对象
func NewTaobaoOmniitemItemPublishRequest() *TaobaoOmniitemItemPublishAPIRequest{
    return &TaobaoOmniitemItemPublishAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniitemItemPublishAPIRequest) GetApiMethodName() string {
    return "taobao.omniitem.item.publish"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniitemItemPublishAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LightPublishInfo Setter
// 发布商品信息
func (r *TaobaoOmniitemItemPublishAPIRequest) SetLightPublishInfo(_lightPublishInfo *ItemLightPublishDTO) error {
    r._lightPublishInfo = _lightPublishInfo
    r.Set("light_publish_info", _lightPublishInfo)
    return nil
}

// LightPublishInfo Getter
func (r TaobaoOmniitemItemPublishAPIRequest) GetLightPublishInfo() *ItemLightPublishDTO {
    return r._lightPublishInfo
}
// OperateType Setter
// 在全域商品或是门店商品中校验码是否重复，可选值对应为ALL或者STORE
func (r *TaobaoOmniitemItemPublishAPIRequest) SetOperateType(_operateType string) error {
    r._operateType = _operateType
    r.Set("operate_type", _operateType)
    return nil
}

// OperateType Getter
func (r TaobaoOmniitemItemPublishAPIRequest) GetOperateType() string {
    return r._operateType
}
