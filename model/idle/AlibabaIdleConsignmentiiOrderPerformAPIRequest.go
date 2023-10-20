package idle

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleConsignmentiiOrderPerformAPIRequest 寄卖V2订单履约 API请求
// alibaba.idle.consignmentii.order.perform
//
// 寄卖V2订单履约，服务商同步订单信息，驱动交易流转
type AlibabaIdleConsignmentiiOrderPerformAPIRequest struct {
	model.Params
	// 同步参数
	_consignmentV2OrderSynDto *ConsignmentV2OrderSynDto
}

// NewAlibabaIdleConsignmentiiOrderPerformRequest 初始化AlibabaIdleConsignmentiiOrderPerformAPIRequest对象
func NewAlibabaIdleConsignmentiiOrderPerformRequest() *AlibabaIdleConsignmentiiOrderPerformAPIRequest {
	return &AlibabaIdleConsignmentiiOrderPerformAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleConsignmentiiOrderPerformAPIRequest) Reset() {
	r._consignmentV2OrderSynDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleConsignmentiiOrderPerformAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.consignmentii.order.perform"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleConsignmentiiOrderPerformAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleConsignmentiiOrderPerformAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetConsignmentV2OrderSynDto is ConsignmentV2OrderSynDto Setter
// 同步参数
func (r *AlibabaIdleConsignmentiiOrderPerformAPIRequest) SetConsignmentV2OrderSynDto(_consignmentV2OrderSynDto *ConsignmentV2OrderSynDto) error {
	r._consignmentV2OrderSynDto = _consignmentV2OrderSynDto
	r.Set("consignment_v2_order_syn_dto", _consignmentV2OrderSynDto)
	return nil
}

// GetConsignmentV2OrderSynDto ConsignmentV2OrderSynDto Getter
func (r AlibabaIdleConsignmentiiOrderPerformAPIRequest) GetConsignmentV2OrderSynDto() *ConsignmentV2OrderSynDto {
	return r._consignmentV2OrderSynDto
}

var poolAlibabaIdleConsignmentiiOrderPerformAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleConsignmentiiOrderPerformRequest()
	},
}

// GetAlibabaIdleConsignmentiiOrderPerformRequest 从 sync.Pool 获取 AlibabaIdleConsignmentiiOrderPerformAPIRequest
func GetAlibabaIdleConsignmentiiOrderPerformAPIRequest() *AlibabaIdleConsignmentiiOrderPerformAPIRequest {
	return poolAlibabaIdleConsignmentiiOrderPerformAPIRequest.Get().(*AlibabaIdleConsignmentiiOrderPerformAPIRequest)
}

// ReleaseAlibabaIdleConsignmentiiOrderPerformAPIRequest 将 AlibabaIdleConsignmentiiOrderPerformAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleConsignmentiiOrderPerformAPIRequest(v *AlibabaIdleConsignmentiiOrderPerformAPIRequest) {
	v.Reset()
	poolAlibabaIdleConsignmentiiOrderPerformAPIRequest.Put(v)
}
