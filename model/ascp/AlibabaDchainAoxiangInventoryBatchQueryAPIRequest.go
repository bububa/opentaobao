package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxianginventorybatchqueryAPIRequest 批量查询库存 API请求
// alibaba.dchain.aoxiang.inventory.batch.query
//
// 批量查询库存
type AlibabadchainaoxianginventorybatchqueryAPIRequest struct {
	model.Params
	// 批量查询库存入参
	_batchQueryInventoryRequest *BatchQueryInventoryRequest
}

// NewAlibabadchainaoxianginventorybatchqueryRequest 初始化AlibabadchainaoxianginventorybatchqueryAPIRequest对象
func NewAlibabadchainaoxianginventorybatchqueryRequest() *AlibabadchainaoxianginventorybatchqueryAPIRequest {
	return &AlibabadchainaoxianginventorybatchqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxianginventorybatchqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.inventory.batch.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxianginventorybatchqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxianginventorybatchqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBatchQueryInventoryRequest is BatchQueryInventoryRequest Setter
// 批量查询库存入参
func (r *AlibabadchainaoxianginventorybatchqueryAPIRequest) SetBatchQueryInventoryRequest(_batchQueryInventoryRequest *BatchQueryInventoryRequest) error {
	r._batchQueryInventoryRequest = _batchQueryInventoryRequest
	r.Set("batch_query_inventory_request", _batchQueryInventoryRequest)
	return nil
}

// GetBatchQueryInventoryRequest BatchQueryInventoryRequest Getter
func (r AlibabadchainaoxianginventorybatchqueryAPIRequest) GetBatchQueryInventoryRequest() *BatchQueryInventoryRequest {
	return r._batchQueryInventoryRequest
}
