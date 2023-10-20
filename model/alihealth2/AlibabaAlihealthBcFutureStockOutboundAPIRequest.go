package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthBcFutureStockOutboundAPIRequest 供应商期货出库 API请求
// alibaba.alihealth.bc.future.stock.outbound
//
// 供应商期货出库
type AlibabaAlihealthBcFutureStockOutboundAPIRequest struct {
	model.Params
	// 请求体
	_futureOutboundReqDto *FutureOutboundReqDto
}

// NewAlibabaAlihealthBcFutureStockOutboundRequest 初始化AlibabaAlihealthBcFutureStockOutboundAPIRequest对象
func NewAlibabaAlihealthBcFutureStockOutboundRequest() *AlibabaAlihealthBcFutureStockOutboundAPIRequest {
	return &AlibabaAlihealthBcFutureStockOutboundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthBcFutureStockOutboundAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.bc.future.stock.outbound"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthBcFutureStockOutboundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthBcFutureStockOutboundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFutureOutboundReqDto is FutureOutboundReqDto Setter
// 请求体
func (r *AlibabaAlihealthBcFutureStockOutboundAPIRequest) SetFutureOutboundReqDto(_futureOutboundReqDto *FutureOutboundReqDto) error {
	r._futureOutboundReqDto = _futureOutboundReqDto
	r.Set("future_outbound_req_dto", _futureOutboundReqDto)
	return nil
}

// GetFutureOutboundReqDto FutureOutboundReqDto Getter
func (r AlibabaAlihealthBcFutureStockOutboundAPIRequest) GetFutureOutboundReqDto() *FutureOutboundReqDto {
	return r._futureOutboundReqDto
}
