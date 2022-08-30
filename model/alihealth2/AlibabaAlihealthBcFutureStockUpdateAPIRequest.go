package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthBcFutureStockUpdateAPIRequest 供应商更新期货库存 API请求
// alibaba.alihealth.bc.future.stock.update
//
// 供应商更新期货库存
type AlibabaAlihealthBcFutureStockUpdateAPIRequest struct {
	model.Params
	// 请求体
	_futureStockReqDto *FutureStockReqDto
}

// NewAlibabaAlihealthBcFutureStockUpdateRequest 初始化AlibabaAlihealthBcFutureStockUpdateAPIRequest对象
func NewAlibabaAlihealthBcFutureStockUpdateRequest() *AlibabaAlihealthBcFutureStockUpdateAPIRequest {
	return &AlibabaAlihealthBcFutureStockUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthBcFutureStockUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.bc.future.stock.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthBcFutureStockUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetFutureStockReqDto is FutureStockReqDto Setter
// 请求体
func (r *AlibabaAlihealthBcFutureStockUpdateAPIRequest) SetFutureStockReqDto(_futureStockReqDto *FutureStockReqDto) error {
	r._futureStockReqDto = _futureStockReqDto
	r.Set("future_stock_req_dto", _futureStockReqDto)
	return nil
}

// GetFutureStockReqDto FutureStockReqDto Getter
func (r AlibabaAlihealthBcFutureStockUpdateAPIRequest) GetFutureStockReqDto() *FutureStockReqDto {
	return r._futureStockReqDto
}
