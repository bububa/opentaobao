package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangConsignorderBatchQueryAPIRequest 发货单批量查询 API请求
// alibaba.dchain.aoxiang.consignorder.batch.query
//
// 发货单批量查询
type AlibabaDchainAoxiangConsignorderBatchQueryAPIRequest struct {
	model.Params
	// 批量查询发货单入参
	_batchQueryConsignorderRequest *BatchQueryConsignorderRequest
}

// NewAlibabaDchainAoxiangConsignorderBatchQueryRequest 初始化AlibabaDchainAoxiangConsignorderBatchQueryAPIRequest对象
func NewAlibabaDchainAoxiangConsignorderBatchQueryRequest() *AlibabaDchainAoxiangConsignorderBatchQueryAPIRequest {
	return &AlibabaDchainAoxiangConsignorderBatchQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangConsignorderBatchQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.consignorder.batch.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangConsignorderBatchQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBatchQueryConsignorderRequest is BatchQueryConsignorderRequest Setter
// 批量查询发货单入参
func (r *AlibabaDchainAoxiangConsignorderBatchQueryAPIRequest) SetBatchQueryConsignorderRequest(_batchQueryConsignorderRequest *BatchQueryConsignorderRequest) error {
	r._batchQueryConsignorderRequest = _batchQueryConsignorderRequest
	r.Set("batch_query_consignorder_request", _batchQueryConsignorderRequest)
	return nil
}

// GetBatchQueryConsignorderRequest BatchQueryConsignorderRequest Getter
func (r AlibabaDchainAoxiangConsignorderBatchQueryAPIRequest) GetBatchQueryConsignorderRequest() *BatchQueryConsignorderRequest {
	return r._batchQueryConsignorderRequest
}
