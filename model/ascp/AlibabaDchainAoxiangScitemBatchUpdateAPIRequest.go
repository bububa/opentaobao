package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangScitemBatchUpdateAPIRequest alibaba.dchain.aoxiang.scitem.batch.update API请求
// alibaba.dchain.aoxiang.scitem.batch.update
//
// 更新货品
type AlibabaDchainAoxiangScitemBatchUpdateAPIRequest struct {
	model.Params
	// 批量更新货品入参，最多30条
	_batchUpdateScitemRequest *BatchUpdateScItemRequest
}

// NewAlibabaDchainAoxiangScitemBatchUpdateRequest 初始化AlibabaDchainAoxiangScitemBatchUpdateAPIRequest对象
func NewAlibabaDchainAoxiangScitemBatchUpdateRequest() *AlibabaDchainAoxiangScitemBatchUpdateAPIRequest {
	return &AlibabaDchainAoxiangScitemBatchUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainAoxiangScitemBatchUpdateAPIRequest) Reset() {
	r._batchUpdateScitemRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangScitemBatchUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.scitem.batch.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangScitemBatchUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangScitemBatchUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBatchUpdateScitemRequest is BatchUpdateScitemRequest Setter
// 批量更新货品入参，最多30条
func (r *AlibabaDchainAoxiangScitemBatchUpdateAPIRequest) SetBatchUpdateScitemRequest(_batchUpdateScitemRequest *BatchUpdateScItemRequest) error {
	r._batchUpdateScitemRequest = _batchUpdateScitemRequest
	r.Set("batch_update_scitem_request", _batchUpdateScitemRequest)
	return nil
}

// GetBatchUpdateScitemRequest BatchUpdateScitemRequest Getter
func (r AlibabaDchainAoxiangScitemBatchUpdateAPIRequest) GetBatchUpdateScitemRequest() *BatchUpdateScItemRequest {
	return r._batchUpdateScitemRequest
}

var poolAlibabaDchainAoxiangScitemBatchUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainAoxiangScitemBatchUpdateRequest()
	},
}

// GetAlibabaDchainAoxiangScitemBatchUpdateRequest 从 sync.Pool 获取 AlibabaDchainAoxiangScitemBatchUpdateAPIRequest
func GetAlibabaDchainAoxiangScitemBatchUpdateAPIRequest() *AlibabaDchainAoxiangScitemBatchUpdateAPIRequest {
	return poolAlibabaDchainAoxiangScitemBatchUpdateAPIRequest.Get().(*AlibabaDchainAoxiangScitemBatchUpdateAPIRequest)
}

// ReleaseAlibabaDchainAoxiangScitemBatchUpdateAPIRequest 将 AlibabaDchainAoxiangScitemBatchUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainAoxiangScitemBatchUpdateAPIRequest(v *AlibabaDchainAoxiangScitemBatchUpdateAPIRequest) {
	v.Reset()
	poolAlibabaDchainAoxiangScitemBatchUpdateAPIRequest.Put(v)
}
