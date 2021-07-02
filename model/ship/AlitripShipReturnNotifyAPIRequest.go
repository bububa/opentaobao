package ship

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripShipReturnNotifyAPIRequest) GetApiMethodName() string {
	return "alitrip.ship.return.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripShipReturnNotifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
