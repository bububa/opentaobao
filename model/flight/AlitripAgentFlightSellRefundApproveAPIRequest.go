package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentFlightSellRefundApproveAPIRequest 销售退票确认 API请求
// alitrip.agent.flight.sell.refund.approve
//
// 销售退票确认
type AlitripAgentFlightSellRefundApproveAPIRequest struct {
	model.Params
	// 入参
	_param *RefundApproveRequestDto
}

// NewAlitripAgentFlightSellRefundApproveRequest 初始化AlitripAgentFlightSellRefundApproveAPIRequest对象
func NewAlitripAgentFlightSellRefundApproveRequest() *AlitripAgentFlightSellRefundApproveAPIRequest {
	return &AlitripAgentFlightSellRefundApproveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripAgentFlightSellRefundApproveAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.flight.sell.refund.approve"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripAgentFlightSellRefundApproveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam is Param Setter
// 入参
func (r *AlitripAgentFlightSellRefundApproveAPIRequest) SetParam(_param *RefundApproveRequestDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlitripAgentFlightSellRefundApproveAPIRequest) GetParam() *RefundApproveRequestDto {
	return r._param
}
