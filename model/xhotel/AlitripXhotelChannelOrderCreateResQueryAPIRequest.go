package xhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripxhotelchannelordercreateresqueryAPIRequest 分销订单查询订单创建结果 API请求
// alitrip.xhotel.channel.order.create.res.query
//
// 针对分销渠道订单，在调用创建订单接口失败1分钟后，调用此接口，用以确认订单是否创建成功。
type AlitripxhotelchannelordercreateresqueryAPIRequest struct {
	model.Params
	// 外部渠道订单号
	_outSourceOrderId string
}

// NewAlitripxhotelchannelordercreateresqueryRequest 初始化AlitripxhotelchannelordercreateresqueryAPIRequest对象
func NewAlitripxhotelchannelordercreateresqueryRequest() *AlitripxhotelchannelordercreateresqueryAPIRequest {
	return &AlitripxhotelchannelordercreateresqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripxhotelchannelordercreateresqueryAPIRequest) GetApiMethodName() string {
	return "alitrip.xhotel.channel.order.create.res.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripxhotelchannelordercreateresqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripxhotelchannelordercreateresqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutSourceOrderId is OutSourceOrderId Setter
// 外部渠道订单号
func (r *AlitripxhotelchannelordercreateresqueryAPIRequest) SetOutSourceOrderId(_outSourceOrderId string) error {
	r._outSourceOrderId = _outSourceOrderId
	r.Set("out_source_order_id", _outSourceOrderId)
	return nil
}

// GetOutSourceOrderId OutSourceOrderId Getter
func (r AlitripxhotelchannelordercreateresqueryAPIRequest) GetOutSourceOrderId() string {
	return r._outSourceOrderId
}
