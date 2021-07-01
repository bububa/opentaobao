package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniitemItemFullupdateAPIRequest
全渠道商品全量更新 API请求
taobao.omniitem.item.fullupdate

全渠道商品全量更新，仅适用于全渠道门店商品
需要全量传入商品相关所有参数，更新时会根据传入的字段进行全量更新
对于SKU信息，会以skus对象进行判断，若传入的skus对象的sku为商品之前未包含的，则新增SKU，如果原先商品有该sku但现在没有传，则删除该SKU。所有传入的SKU信息要么全部均传入skuId，要么全部都不传入skuId。对于新增SKU的场景，目前无需传入SKUID，会根据传入的销售属性自动对应 */
type TaobaoOmniitemItemFullupdateAPIRequest struct {
	model.Params
	// 发布商品信息
	_lightPublishInfo *ItemLightPublishDto
	// 操作类型，STORE表示门店域新增，ALL表示全域新增
	_operateType string
}

// NewTaobaoOmniitemItemFullupdateRequest 初始化TaobaoOmniitemItemFullupdateAPIRequest对象
func NewTaobaoOmniitemItemFullupdateRequest() *TaobaoOmniitemItemFullupdateAPIRequest {
	return &TaobaoOmniitemItemFullupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniitemItemFullupdateAPIRequest) GetApiMethodName() string {
	return "taobao.omniitem.item.fullupdate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniitemItemFullupdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is LightPublishInfo Setter
// 发布商品信息
func (r *TaobaoOmniitemItemFullupdateAPIRequest) SetLightPublishInfo(_lightPublishInfo *ItemLightPublishDto) error {
	r._lightPublishInfo = _lightPublishInfo
	r.Set("light_publish_info", _lightPublishInfo)
	return nil
}

// Get LightPublishInfo Getter
func (r TaobaoOmniitemItemFullupdateAPIRequest) GetLightPublishInfo() *ItemLightPublishDto {
	return r._lightPublishInfo
}

// Set is OperateType Setter
// 操作类型，STORE表示门店域新增，ALL表示全域新增
func (r *TaobaoOmniitemItemFullupdateAPIRequest) SetOperateType(_operateType string) error {
	r._operateType = _operateType
	r.Set("operate_type", _operateType)
	return nil
}

// Get OperateType Getter
func (r TaobaoOmniitemItemFullupdateAPIRequest) GetOperateType() string {
	return r._operateType
}
