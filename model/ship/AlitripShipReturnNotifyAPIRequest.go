package ship

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripShipReturnNotifyAPIRequest 船票退票退款回填接口 API请求
// alitrip.ship.return.notify
//
// 此接口为接入商调用飞猪接口回填退票状态，飞猪平台给用户进行退票退款。飞猪平台保证数据幂等。
type AlitripShipReturnNotifyAPIRequest struct {
	model.Params
	// 退票请求入参
	_confirmRefundRQ *ShipAgentConfirmRefundRq
}

// NewAlitripShipReturnNotifyRequest 初始化AlitripShipReturnNotifyAPIRequest对象
func NewAlitripShipReturnNotifyRequest() *AlitripShipReturnNotifyAPIRequest {
	return &AlitripShipReturnNotifyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripShipReturnNotifyAPIRequest) Reset() {
	r._confirmRefundRQ = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripShipReturnNotifyAPIRequest) GetApiMethodName() string {
	return "alitrip.ship.return.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripShipReturnNotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripShipReturnNotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetConfirmRefundRQ is ConfirmRefundRQ Setter
// 退票请求入参
func (r *AlitripShipReturnNotifyAPIRequest) SetConfirmRefundRQ(_confirmRefundRQ *ShipAgentConfirmRefundRq) error {
	r._confirmRefundRQ = _confirmRefundRQ
	r.Set("confirm_refund_r_q", _confirmRefundRQ)
	return nil
}

// GetConfirmRefundRQ ConfirmRefundRQ Getter
func (r AlitripShipReturnNotifyAPIRequest) GetConfirmRefundRQ() *ShipAgentConfirmRefundRq {
	return r._confirmRefundRQ
}

var poolAlitripShipReturnNotifyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripShipReturnNotifyRequest()
	},
}

// GetAlitripShipReturnNotifyRequest 从 sync.Pool 获取 AlitripShipReturnNotifyAPIRequest
func GetAlitripShipReturnNotifyAPIRequest() *AlitripShipReturnNotifyAPIRequest {
	return poolAlitripShipReturnNotifyAPIRequest.Get().(*AlitripShipReturnNotifyAPIRequest)
}

// ReleaseAlitripShipReturnNotifyAPIRequest 将 AlitripShipReturnNotifyAPIRequest 放入 sync.Pool
func ReleaseAlitripShipReturnNotifyAPIRequest(v *AlitripShipReturnNotifyAPIRequest) {
	v.Reset()
	poolAlitripShipReturnNotifyAPIRequest.Put(v)
}
