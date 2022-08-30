package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIRequest 组合货品新建&更新 API请求
// alibaba.dchain.aoxiang.combineitem.batch.update.async
//
// 组合货品新建&更新
type AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIRequest struct {
	model.Params
	// 请求入参
	_combineItemBatchUpsertAsyncRequest *CombineItemBatchUpsertAsyncRequest
}

// NewAlibabaDchainAoxiangCombineitemBatchUpdateAsyncRequest 初始化AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIRequest对象
func NewAlibabaDchainAoxiangCombineitemBatchUpdateAsyncRequest() *AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIRequest {
	return &AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.combineitem.batch.update.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
