package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleConsignmentOrderPerformAPIRequest
帮卖订单履约 API请求
alibaba.idle.consignment.order.perform

帮卖订单履约，回收商同步订单信息，驱动交易流转 */
type AlibabaIdleConsignmentOrderPerformAPIRequest struct {
	model.Params
	// 帮卖订单同步DTO
	_param *ConsignmentOrderSynDto
}

// NewAlibabaIdleConsignmentOrderPerformRequest 初始化AlibabaIdleConsignmentOrderPerformAPIRequest对象
func NewAlibabaIdleConsignmentOrderPerformRequest() *AlibabaIdleConsignmentOrderPerformAPIRequest {
	return &AlibabaIdleConsignmentOrderPerformAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleConsignmentOrderPerformAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.consignment.order.perform"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleConsignmentOrderPerformAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param Setter
// 帮卖订单同步DTO
func (r *AlibabaIdleConsignmentOrderPerformAPIRequest) SetParam(_param *ConsignmentOrderSynDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r AlibabaIdleConsignmentOrderPerformAPIRequest) GetParam() *ConsignmentOrderSynDto {
	return r._param
}
