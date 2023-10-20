package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripxhotelchannelordercreateAPIRequest 渠道分销创建订单接口 API请求
// alitrip.xhotel.channel.order.create
//
// 创建订单接口服务（如菲住等其他渠道分销提供）
type AlitripxhotelchannelordercreateAPIRequest struct {
	model.Params
	// 创建订单参数
	_outSourceOrderCreateReq *OutSourceOrderCreateReq
}

// NewAlitripxhotelchannelordercreateRequest 初始化AlitripxhotelchannelordercreateAPIRequest对象
func NewAlitripxhotelchannelordercreateRequest() *AlitripxhotelchannelordercreateAPIRequest {
	return &AlitripxhotelchannelordercreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripxhotelchannelordercreateAPIRequest) GetApiMethodName() string {
	return "alitrip.xhotel.channel.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripxhotelchannelordercreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripxhotelchannelordercreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutSourceOrderCreateReq is OutSourceOrderCreateReq Setter
// 创建订单参数
func (r *AlitripxhotelchannelordercreateAPIRequest) SetOutSourceOrderCreateReq(_outSourceOrderCreateReq *OutSourceOrderCreateReq) error {
	r._outSourceOrderCreateReq = _outSourceOrderCreateReq
	r.Set("out_source_order_create_req", _outSourceOrderCreateReq)
	return nil
}

// GetOutSourceOrderCreateReq OutSourceOrderCreateReq Getter
func (r AlitripxhotelchannelordercreateAPIRequest) GetOutSourceOrderCreateReq() *OutSourceOrderCreateReq {
	return r._outSourceOrderCreateReq
}
