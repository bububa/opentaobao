package trade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAscpOrdersSaleCreateAPIRequest ASCP渠道中心销售单创建接口 API请求
// tmall.ascp.orders.sale.create
//
// ASCP渠道中心销售单创建接口
type TmallAscpOrdersSaleCreateAPIRequest struct {
	model.Params
	// 请求对象
	_channelOrderRequest *CreateChannelOrderRequest
}

// NewTmallAscpOrdersSaleCreateRequest 初始化TmallAscpOrdersSaleCreateAPIRequest对象
func NewTmallAscpOrdersSaleCreateRequest() *TmallAscpOrdersSaleCreateAPIRequest {
	return &TmallAscpOrdersSaleCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallAscpOrdersSaleCreateAPIRequest) Reset() {
	r._channelOrderRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAscpOrdersSaleCreateAPIRequest) GetApiMethodName() string {
	return "tmall.ascp.orders.sale.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAscpOrdersSaleCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallAscpOrdersSaleCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChannelOrderRequest is ChannelOrderRequest Setter
// 请求对象
func (r *TmallAscpOrdersSaleCreateAPIRequest) SetChannelOrderRequest(_channelOrderRequest *CreateChannelOrderRequest) error {
	r._channelOrderRequest = _channelOrderRequest
	r.Set("channel_order_request", _channelOrderRequest)
	return nil
}

// GetChannelOrderRequest ChannelOrderRequest Getter
func (r TmallAscpOrdersSaleCreateAPIRequest) GetChannelOrderRequest() *CreateChannelOrderRequest {
	return r._channelOrderRequest
}

var poolTmallAscpOrdersSaleCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallAscpOrdersSaleCreateRequest()
	},
}

// GetTmallAscpOrdersSaleCreateRequest 从 sync.Pool 获取 TmallAscpOrdersSaleCreateAPIRequest
func GetTmallAscpOrdersSaleCreateAPIRequest() *TmallAscpOrdersSaleCreateAPIRequest {
	return poolTmallAscpOrdersSaleCreateAPIRequest.Get().(*TmallAscpOrdersSaleCreateAPIRequest)
}

// ReleaseTmallAscpOrdersSaleCreateAPIRequest 将 TmallAscpOrdersSaleCreateAPIRequest 放入 sync.Pool
func ReleaseTmallAscpOrdersSaleCreateAPIRequest(v *TmallAscpOrdersSaleCreateAPIRequest) {
	v.Reset()
	poolTmallAscpOrdersSaleCreateAPIRequest.Put(v)
}
