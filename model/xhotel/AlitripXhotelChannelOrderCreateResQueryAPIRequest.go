package xhotel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripXhotelChannelOrderCreateResQueryAPIRequest 分销订单查询订单创建结果 API请求
// alitrip.xhotel.channel.order.create.res.query
//
// 针对分销渠道订单，在调用创建订单接口失败1分钟后，调用此接口，用以确认订单是否创建成功。
type AlitripXhotelChannelOrderCreateResQueryAPIRequest struct {
	model.Params
	// 外部渠道订单号
	_outSourceOrderId string
}

// NewAlitripXhotelChannelOrderCreateResQueryRequest 初始化AlitripXhotelChannelOrderCreateResQueryAPIRequest对象
func NewAlitripXhotelChannelOrderCreateResQueryRequest() *AlitripXhotelChannelOrderCreateResQueryAPIRequest {
	return &AlitripXhotelChannelOrderCreateResQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripXhotelChannelOrderCreateResQueryAPIRequest) Reset() {
	r._outSourceOrderId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripXhotelChannelOrderCreateResQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.xhotel.channel.order.create.res.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripXhotelChannelOrderCreateResQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripXhotelChannelOrderCreateResQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutSourceOrderId is OutSourceOrderId Setter
// 外部渠道订单号
func (r *AlitripXhotelChannelOrderCreateResQueryAPIRequest) SetOutSourceOrderId(_outSourceOrderId string) error {
	r._outSourceOrderId = _outSourceOrderId
	r.Set("out_source_order_id", _outSourceOrderId)
	return nil
}

// GetOutSourceOrderId OutSourceOrderId Getter
func (r AlitripXhotelChannelOrderCreateResQueryAPIRequest) GetOutSourceOrderId() string {
	return r._outSourceOrderId
}

var poolAlitripXhotelChannelOrderCreateResQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripXhotelChannelOrderCreateResQueryRequest()
	},
}

// GetAlitripXhotelChannelOrderCreateResQueryRequest 从 sync.Pool 获取 AlitripXhotelChannelOrderCreateResQueryAPIRequest
func GetAlitripXhotelChannelOrderCreateResQueryAPIRequest() *AlitripXhotelChannelOrderCreateResQueryAPIRequest {
	return poolAlitripXhotelChannelOrderCreateResQueryAPIRequest.Get().(*AlitripXhotelChannelOrderCreateResQueryAPIRequest)
}

// ReleaseAlitripXhotelChannelOrderCreateResQueryAPIRequest 将 AlitripXhotelChannelOrderCreateResQueryAPIRequest 放入 sync.Pool
func ReleaseAlitripXhotelChannelOrderCreateResQueryAPIRequest(v *AlitripXhotelChannelOrderCreateResQueryAPIRequest) {
	v.Reset()
	poolAlitripXhotelChannelOrderCreateResQueryAPIRequest.Put(v)
}
