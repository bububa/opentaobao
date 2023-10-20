package ascp

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainAoxiangConsignorderBatchQueryAPIRequest) Reset() {
	r._batchQueryConsignorderRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangConsignorderBatchQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.consignorder.batch.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangConsignorderBatchQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangConsignorderBatchQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaDchainAoxiangConsignorderBatchQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainAoxiangConsignorderBatchQueryRequest()
	},
}

// GetAlibabaDchainAoxiangConsignorderBatchQueryRequest 从 sync.Pool 获取 AlibabaDchainAoxiangConsignorderBatchQueryAPIRequest
func GetAlibabaDchainAoxiangConsignorderBatchQueryAPIRequest() *AlibabaDchainAoxiangConsignorderBatchQueryAPIRequest {
	return poolAlibabaDchainAoxiangConsignorderBatchQueryAPIRequest.Get().(*AlibabaDchainAoxiangConsignorderBatchQueryAPIRequest)
}

// ReleaseAlibabaDchainAoxiangConsignorderBatchQueryAPIRequest 将 AlibabaDchainAoxiangConsignorderBatchQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainAoxiangConsignorderBatchQueryAPIRequest(v *AlibabaDchainAoxiangConsignorderBatchQueryAPIRequest) {
	v.Reset()
	poolAlibabaDchainAoxiangConsignorderBatchQueryAPIRequest.Put(v)
}
