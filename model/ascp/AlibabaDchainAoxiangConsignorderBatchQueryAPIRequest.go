package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangconsignorderbatchqueryAPIRequest 发货单批量查询 API请求
// alibaba.dchain.aoxiang.consignorder.batch.query
//
// 发货单批量查询
type AlibabadchainaoxiangconsignorderbatchqueryAPIRequest struct {
	model.Params
	// 批量查询发货单入参
	_batchQueryConsignorderRequest *BatchQueryConsignorderRequest
}

// NewAlibabadchainaoxiangconsignorderbatchqueryRequest 初始化AlibabadchainaoxiangconsignorderbatchqueryAPIRequest对象
func NewAlibabadchainaoxiangconsignorderbatchqueryRequest() *AlibabadchainaoxiangconsignorderbatchqueryAPIRequest {
	return &AlibabadchainaoxiangconsignorderbatchqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangconsignorderbatchqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.consignorder.batch.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangconsignorderbatchqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangconsignorderbatchqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBatchQueryConsignorderRequest is BatchQueryConsignorderRequest Setter
// 批量查询发货单入参
func (r *AlibabadchainaoxiangconsignorderbatchqueryAPIRequest) SetBatchQueryConsignorderRequest(_batchQueryConsignorderRequest *BatchQueryConsignorderRequest) error {
	r._batchQueryConsignorderRequest = _batchQueryConsignorderRequest
	r.Set("batch_query_consignorder_request", _batchQueryConsignorderRequest)
	return nil
}

// GetBatchQueryConsignorderRequest BatchQueryConsignorderRequest Getter
func (r AlibabadchainaoxiangconsignorderbatchqueryAPIRequest) GetBatchQueryConsignorderRequest() *BatchQueryConsignorderRequest {
	return r._batchQueryConsignorderRequest
}
