package miniapp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIRequest ISV查询应用的渠道信息 API请求
// taobao.miniapp.ext.delivery.app.channel.configs.query
//
// ISV查询应用的渠道信息
type TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIRequest struct {
	model.Params
	// 入参
	_req *AppChannelQueryRequest
}

// NewTaobaoMiniappExtDeliveryAppChannelConfigsQueryRequest 初始化TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIRequest对象
func NewTaobaoMiniappExtDeliveryAppChannelConfigsQueryRequest() *TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIRequest {
	return &TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.ext.delivery.app.channel.configs.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReq is Req Setter
// 入参
func (r *TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIRequest) SetReq(_req *AppChannelQueryRequest) error {
	r._req = _req
	r.Set("req", _req)
	return nil
}

// GetReq Req Getter
func (r TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIRequest) GetReq() *AppChannelQueryRequest {
	return r._req
}
