package idle

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleConsignmentOrderPerformAPIRequest 帮卖订单履约 API请求
// alibaba.idle.consignment.order.perform
//
// 帮卖订单履约，回收商同步订单信息，驱动交易流转
type AlibabaIdleConsignmentOrderPerformAPIRequest struct {
	model.Params
	// 帮卖订单同步DTO
	_param *ConsignmentOrderSynDto
}

// NewAlibabaIdleConsignmentOrderPerformRequest 初始化AlibabaIdleConsignmentOrderPerformAPIRequest对象
func NewAlibabaIdleConsignmentOrderPerformRequest() *AlibabaIdleConsignmentOrderPerformAPIRequest {
	return &AlibabaIdleConsignmentOrderPerformAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleConsignmentOrderPerformAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleConsignmentOrderPerformAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.consignment.order.perform"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleConsignmentOrderPerformAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleConsignmentOrderPerformAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 帮卖订单同步DTO
func (r *AlibabaIdleConsignmentOrderPerformAPIRequest) SetParam(_param *ConsignmentOrderSynDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaIdleConsignmentOrderPerformAPIRequest) GetParam() *ConsignmentOrderSynDto {
	return r._param
}

var poolAlibabaIdleConsignmentOrderPerformAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleConsignmentOrderPerformRequest()
	},
}

// GetAlibabaIdleConsignmentOrderPerformRequest 从 sync.Pool 获取 AlibabaIdleConsignmentOrderPerformAPIRequest
func GetAlibabaIdleConsignmentOrderPerformAPIRequest() *AlibabaIdleConsignmentOrderPerformAPIRequest {
	return poolAlibabaIdleConsignmentOrderPerformAPIRequest.Get().(*AlibabaIdleConsignmentOrderPerformAPIRequest)
}

// ReleaseAlibabaIdleConsignmentOrderPerformAPIRequest 将 AlibabaIdleConsignmentOrderPerformAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleConsignmentOrderPerformAPIRequest(v *AlibabaIdleConsignmentOrderPerformAPIRequest) {
	v.Reset()
	poolAlibabaIdleConsignmentOrderPerformAPIRequest.Put(v)
}
