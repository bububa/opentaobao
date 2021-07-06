package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniitemItemPublishAPIRequest 全渠道门店商品轻发布 API请求
// taobao.omniitem.item.publish
//
// 全渠道门店商品轻发布
type TaobaoOmniitemItemPublishAPIRequest struct {
	model.Params
	// 在全域商品或是门店商品中校验码是否重复，可选值对应为ALL或者STORE
	_operateType string
	// 发布商品信息
	_lightPublishInfo *ItemLightPublishDto
}

// NewTaobaoOmniitemItemPublishRequest 初始化TaobaoOmniitemItemPublishAPIRequest对象
func NewTaobaoOmniitemItemPublishRequest() *TaobaoOmniitemItemPublishAPIRequest {
	return &TaobaoOmniitemItemPublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniitemItemPublishAPIRequest) GetApiMethodName() string {
	return "taobao.omniitem.item.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniitemItemPublishAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOperateType is OperateType Setter
// 在全域商品或是门店商品中校验码是否重复，可选值对应为ALL或者STORE
func (r *TaobaoOmniitemItemPublishAPIRequest) SetOperateType(_operateType string) error {
	r._operateType = _operateType
	r.Set("operate_type", _operateType)
	return nil
}

// GetOperateType OperateType Getter
func (r TaobaoOmniitemItemPublishAPIRequest) GetOperateType() string {
	return r._operateType
}

// SetLightPublishInfo is LightPublishInfo Setter
// 发布商品信息
func (r *TaobaoOmniitemItemPublishAPIRequest) SetLightPublishInfo(_lightPublishInfo *ItemLightPublishDto) error {
	r._lightPublishInfo = _lightPublishInfo
	r.Set("light_publish_info", _lightPublishInfo)
	return nil
}

// GetLightPublishInfo LightPublishInfo Getter
func (r TaobaoOmniitemItemPublishAPIRequest) GetLightPublishInfo() *ItemLightPublishDto {
	return r._lightPublishInfo
}
