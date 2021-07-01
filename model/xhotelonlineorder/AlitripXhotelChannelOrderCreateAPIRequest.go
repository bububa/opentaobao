package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripXhotelChannelOrderCreateAPIRequest
渠道分销创建订单接口 API请求
alitrip.xhotel.channel.order.create

创建订单接口服务（如菲住等其他渠道分销提供） */
type AlitripXhotelChannelOrderCreateAPIRequest struct {
	model.Params
	// 创建订单参数
	_outSourceOrderCreateReq *OutSourceOrderCreateReq
}

// NewAlitripXhotelChannelOrderCreateRequest 初始化AlitripXhotelChannelOrderCreateAPIRequest对象
func NewAlitripXhotelChannelOrderCreateRequest() *AlitripXhotelChannelOrderCreateAPIRequest {
	return &AlitripXhotelChannelOrderCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripXhotelChannelOrderCreateAPIRequest) GetApiMethodName() string {
	return "alitrip.xhotel.channel.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripXhotelChannelOrderCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OutSourceOrderCreateReq Setter
// 创建订单参数
func (r *AlitripXhotelChannelOrderCreateAPIRequest) SetOutSourceOrderCreateReq(_outSourceOrderCreateReq *OutSourceOrderCreateReq) error {
	r._outSourceOrderCreateReq = _outSourceOrderCreateReq
	r.Set("out_source_order_create_req", _outSourceOrderCreateReq)
	return nil
}

// Get OutSourceOrderCreateReq Getter
func (r AlitripXhotelChannelOrderCreateAPIRequest) GetOutSourceOrderCreateReq() *OutSourceOrderCreateReq {
	return r._outSourceOrderCreateReq
}
