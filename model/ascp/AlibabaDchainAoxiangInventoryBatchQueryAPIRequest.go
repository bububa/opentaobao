package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangInventoryBatchQueryAPIRequest 批量查询库存 API请求
// alibaba.dchain.aoxiang.inventory.batch.query
//
// 批量查询库存
type AlibabaDchainAoxiangInventoryBatchQueryAPIRequest struct {
	model.Params
	// 批量查询库存入参
	_batchQueryInventoryRequest *BatchQueryInventoryRequest
}

// NewAlibabaDchainAoxiangInventoryBatchQueryRequest 初始化AlibabaDchainAoxiangInventoryBatchQueryAPIRequest对象
func NewAlibabaDchainAoxiangInventoryBatchQueryRequest() *AlibabaDchainAoxiangInventoryBatchQueryAPIRequest {
	return &AlibabaDchainAoxiangInventoryBatchQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainAoxiangInventoryBatchQueryAPIRequest) Reset() {
	r._batchQueryInventoryRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangInventoryBatchQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.inventory.batch.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangInventoryBatchQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangInventoryBatchQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBatchQueryInventoryRequest is BatchQueryInventoryRequest Setter
// 批量查询库存入参
func (r *AlibabaDchainAoxiangInventoryBatchQueryAPIRequest) SetBatchQueryInventoryRequest(_batchQueryInventoryRequest *BatchQueryInventoryRequest) error {
	r._batchQueryInventoryRequest = _batchQueryInventoryRequest
	r.Set("batch_query_inventory_request", _batchQueryInventoryRequest)
	return nil
}

// GetBatchQueryInventoryRequest BatchQueryInventoryRequest Getter
func (r AlibabaDchainAoxiangInventoryBatchQueryAPIRequest) GetBatchQueryInventoryRequest() *BatchQueryInventoryRequest {
	return r._batchQueryInventoryRequest
}

var poolAlibabaDchainAoxiangInventoryBatchQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainAoxiangInventoryBatchQueryRequest()
	},
}

// GetAlibabaDchainAoxiangInventoryBatchQueryRequest 从 sync.Pool 获取 AlibabaDchainAoxiangInventoryBatchQueryAPIRequest
func GetAlibabaDchainAoxiangInventoryBatchQueryAPIRequest() *AlibabaDchainAoxiangInventoryBatchQueryAPIRequest {
	return poolAlibabaDchainAoxiangInventoryBatchQueryAPIRequest.Get().(*AlibabaDchainAoxiangInventoryBatchQueryAPIRequest)
}

// ReleaseAlibabaDchainAoxiangInventoryBatchQueryAPIRequest 将 AlibabaDchainAoxiangInventoryBatchQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainAoxiangInventoryBatchQueryAPIRequest(v *AlibabaDchainAoxiangInventoryBatchQueryAPIRequest) {
	v.Reset()
	poolAlibabaDchainAoxiangInventoryBatchQueryAPIRequest.Put(v)
}
