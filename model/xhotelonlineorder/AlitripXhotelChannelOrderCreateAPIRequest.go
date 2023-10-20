package xhotelonlineorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripXhotelChannelOrderCreateAPIRequest 渠道分销创建订单接口 API请求
// alitrip.xhotel.channel.order.create
//
// 创建订单接口服务（如菲住等其他渠道分销提供）
type AlitripXhotelChannelOrderCreateAPIRequest struct {
	model.Params
	// 创建订单参数
	_outSourceOrderCreateReq *OutSourceOrderCreateReq
}

// NewAlitripXhotelChannelOrderCreateRequest 初始化AlitripXhotelChannelOrderCreateAPIRequest对象
func NewAlitripXhotelChannelOrderCreateRequest() *AlitripXhotelChannelOrderCreateAPIRequest {
	return &AlitripXhotelChannelOrderCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripXhotelChannelOrderCreateAPIRequest) Reset() {
	r._outSourceOrderCreateReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripXhotelChannelOrderCreateAPIRequest) GetApiMethodName() string {
	return "alitrip.xhotel.channel.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripXhotelChannelOrderCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripXhotelChannelOrderCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutSourceOrderCreateReq is OutSourceOrderCreateReq Setter
// 创建订单参数
func (r *AlitripXhotelChannelOrderCreateAPIRequest) SetOutSourceOrderCreateReq(_outSourceOrderCreateReq *OutSourceOrderCreateReq) error {
	r._outSourceOrderCreateReq = _outSourceOrderCreateReq
	r.Set("out_source_order_create_req", _outSourceOrderCreateReq)
	return nil
}

// GetOutSourceOrderCreateReq OutSourceOrderCreateReq Getter
func (r AlitripXhotelChannelOrderCreateAPIRequest) GetOutSourceOrderCreateReq() *OutSourceOrderCreateReq {
	return r._outSourceOrderCreateReq
}

var poolAlitripXhotelChannelOrderCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripXhotelChannelOrderCreateRequest()
	},
}

// GetAlitripXhotelChannelOrderCreateRequest 从 sync.Pool 获取 AlitripXhotelChannelOrderCreateAPIRequest
func GetAlitripXhotelChannelOrderCreateAPIRequest() *AlitripXhotelChannelOrderCreateAPIRequest {
	return poolAlitripXhotelChannelOrderCreateAPIRequest.Get().(*AlitripXhotelChannelOrderCreateAPIRequest)
}

// ReleaseAlitripXhotelChannelOrderCreateAPIRequest 将 AlitripXhotelChannelOrderCreateAPIRequest 放入 sync.Pool
func ReleaseAlitripXhotelChannelOrderCreateAPIRequest(v *AlitripXhotelChannelOrderCreateAPIRequest) {
	v.Reset()
	poolAlitripXhotelChannelOrderCreateAPIRequest.Put(v)
}
