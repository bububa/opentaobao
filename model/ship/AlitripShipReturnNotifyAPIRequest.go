package ship

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripshipreturnnotifyAPIRequest 船票退票退款回填接口 API请求
// alitrip.ship.return.notify
//
// 此接口为接入商调用飞猪接口回填退票状态，飞猪平台给用户进行退票退款。飞猪平台保证数据幂等。
type AlitripshipreturnnotifyAPIRequest struct {
	model.Params
	// 退票请求入参
	_confirmRefundRQ *ShipAgentConfirmRefundRq
}

// NewAlitripshipreturnnotifyRequest 初始化AlitripshipreturnnotifyAPIRequest对象
func NewAlitripshipreturnnotifyRequest() *AlitripshipreturnnotifyAPIRequest {
	return &AlitripshipreturnnotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripshipreturnnotifyAPIRequest) GetApiMethodName() string {
	return "alitrip.ship.return.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripshipreturnnotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripshipreturnnotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetConfirmRefundRQ is ConfirmRefundRQ Setter
// 退票请求入参
func (r *AlitripshipreturnnotifyAPIRequest) SetConfirmRefundRQ(_confirmRefundRQ *ShipAgentConfirmRefundRq) error {
	r._confirmRefundRQ = _confirmRefundRQ
	r.Set("confirm_refund_r_q", _confirmRefundRQ)
	return nil
}

// GetConfirmRefundRQ ConfirmRefundRQ Getter
func (r AlitripshipreturnnotifyAPIRequest) GetConfirmRefundRQ() *ShipAgentConfirmRefundRq {
	return r._confirmRefundRQ
}
