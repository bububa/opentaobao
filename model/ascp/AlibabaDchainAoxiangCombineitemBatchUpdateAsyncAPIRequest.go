package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIRequest 组合货品新建&更新 API请求
// alibaba.dchain.aoxiang.combineitem.batch.update.async
//
// 组合货品新建&amp;更新
type AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIRequest struct {
	model.Params
	// 请求入参
	_combineItemBatchUpsertAsyncRequest *CombineItemBatchUpsertAsyncRequest
}

// NewAlibabaDchainAoxiangCombineitemBatchUpdateAsyncRequest 初始化AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIRequest对象
func NewAlibabaDchainAoxiangCombineitemBatchUpdateAsyncRequest() *AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIRequest {
	return &AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIRequest) Reset() {
	r._combineItemBatchUpsertAsyncRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.combineitem.batch.update.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCombineItemBatchUpsertAsyncRequest is CombineItemBatchUpsertAsyncRequest Setter
// 请求入参
func (r *AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIRequest) SetCombineItemBatchUpsertAsyncRequest(_combineItemBatchUpsertAsyncRequest *CombineItemBatchUpsertAsyncRequest) error {
	r._combineItemBatchUpsertAsyncRequest = _combineItemBatchUpsertAsyncRequest
	r.Set("combine_item_batch_upsert_async_request", _combineItemBatchUpsertAsyncRequest)
	return nil
}

// GetCombineItemBatchUpsertAsyncRequest CombineItemBatchUpsertAsyncRequest Getter
func (r AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIRequest) GetCombineItemBatchUpsertAsyncRequest() *CombineItemBatchUpsertAsyncRequest {
	return r._combineItemBatchUpsertAsyncRequest
}

var poolAlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainAoxiangCombineitemBatchUpdateAsyncRequest()
	},
}

// GetAlibabaDchainAoxiangCombineitemBatchUpdateAsyncRequest 从 sync.Pool 获取 AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIRequest
func GetAlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIRequest() *AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIRequest {
	return poolAlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIRequest.Get().(*AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIRequest)
}

// ReleaseAlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIRequest 将 AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIRequest(v *AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIRequest) {
	v.Reset()
	poolAlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIRequest.Put(v)
}
