package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleconsignmentorderperformAPIRequest 帮卖订单履约 API请求
// alibaba.idle.consignment.order.perform
//
// 帮卖订单履约，回收商同步订单信息，驱动交易流转
type AlibabaidleconsignmentorderperformAPIRequest struct {
	model.Params
	// 帮卖订单同步DTO
	_param *ConsignmentOrderSynDto
}

// NewAlibabaidleconsignmentorderperformRequest 初始化AlibabaidleconsignmentorderperformAPIRequest对象
func NewAlibabaidleconsignmentorderperformRequest() *AlibabaidleconsignmentorderperformAPIRequest {
	return &AlibabaidleconsignmentorderperformAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleconsignmentorderperformAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.consignment.order.perform"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleconsignmentorderperformAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleconsignmentorderperformAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 帮卖订单同步DTO
func (r *AlibabaidleconsignmentorderperformAPIRequest) SetParam(_param *ConsignmentOrderSynDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaidleconsignmentorderperformAPIRequest) GetParam() *ConsignmentOrderSynDto {
	return r._param
}
