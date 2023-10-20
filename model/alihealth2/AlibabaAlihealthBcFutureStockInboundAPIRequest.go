package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthBcFutureStockInboundAPIRequest 供应商上报期货库存 API请求
// alibaba.alihealth.bc.future.stock.inbound
//
// 供应商上报期货库存
type AlibabaAlihealthBcFutureStockInboundAPIRequest struct {
	model.Params
	// 请求体
	_futureInboundReqDto *FutureInboundReqDto
}

// NewAlibabaAlihealthBcFutureStockInboundRequest 初始化AlibabaAlihealthBcFutureStockInboundAPIRequest对象
func NewAlibabaAlihealthBcFutureStockInboundRequest() *AlibabaAlihealthBcFutureStockInboundAPIRequest {
	return &AlibabaAlihealthBcFutureStockInboundAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthBcFutureStockInboundAPIRequest) Reset() {
	r._futureInboundReqDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthBcFutureStockInboundAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.bc.future.stock.inbound"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthBcFutureStockInboundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthBcFutureStockInboundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFutureInboundReqDto is FutureInboundReqDto Setter
// 请求体
func (r *AlibabaAlihealthBcFutureStockInboundAPIRequest) SetFutureInboundReqDto(_futureInboundReqDto *FutureInboundReqDto) error {
	r._futureInboundReqDto = _futureInboundReqDto
	r.Set("future_inbound_req_dto", _futureInboundReqDto)
	return nil
}

// GetFutureInboundReqDto FutureInboundReqDto Getter
func (r AlibabaAlihealthBcFutureStockInboundAPIRequest) GetFutureInboundReqDto() *FutureInboundReqDto {
	return r._futureInboundReqDto
}

var poolAlibabaAlihealthBcFutureStockInboundAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthBcFutureStockInboundRequest()
	},
}

// GetAlibabaAlihealthBcFutureStockInboundRequest 从 sync.Pool 获取 AlibabaAlihealthBcFutureStockInboundAPIRequest
func GetAlibabaAlihealthBcFutureStockInboundAPIRequest() *AlibabaAlihealthBcFutureStockInboundAPIRequest {
	return poolAlibabaAlihealthBcFutureStockInboundAPIRequest.Get().(*AlibabaAlihealthBcFutureStockInboundAPIRequest)
}

// ReleaseAlibabaAlihealthBcFutureStockInboundAPIRequest 将 AlibabaAlihealthBcFutureStockInboundAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthBcFutureStockInboundAPIRequest(v *AlibabaAlihealthBcFutureStockInboundAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthBcFutureStockInboundAPIRequest.Put(v)
}
