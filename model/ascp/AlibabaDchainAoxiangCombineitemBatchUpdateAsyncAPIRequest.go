package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangcombineitembatchupdateasyncAPIRequest 组合货品新建&更新 API请求
// alibaba.dchain.aoxiang.combineitem.batch.update.async
//
// 组合货品新建&amp;更新
type AlibabadchainaoxiangcombineitembatchupdateasyncAPIRequest struct {
	model.Params
	// 请求入参
	_combineItemBatchUpsertAsyncRequest *CombineItemBatchUpsertAsyncRequest
}

// NewAlibabadchainaoxiangcombineitembatchupdateasyncRequest 初始化AlibabadchainaoxiangcombineitembatchupdateasyncAPIRequest对象
func NewAlibabadchainaoxiangcombineitembatchupdateasyncRequest() *AlibabadchainaoxiangcombineitembatchupdateasyncAPIRequest {
	return &AlibabadchainaoxiangcombineitembatchupdateasyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangcombineitembatchupdateasyncAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.combineitem.batch.update.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangcombineitembatchupdateasyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangcombineitembatchupdateasyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCombineItemBatchUpsertAsyncRequest is CombineItemBatchUpsertAsyncRequest Setter
// 请求入参
func (r *AlibabadchainaoxiangcombineitembatchupdateasyncAPIRequest) SetCombineItemBatchUpsertAsyncRequest(_combineItemBatchUpsertAsyncRequest *CombineItemBatchUpsertAsyncRequest) error {
	r._combineItemBatchUpsertAsyncRequest = _combineItemBatchUpsertAsyncRequest
	r.Set("combine_item_batch_upsert_async_request", _combineItemBatchUpsertAsyncRequest)
	return nil
}

// GetCombineItemBatchUpsertAsyncRequest CombineItemBatchUpsertAsyncRequest Getter
func (r AlibabadchainaoxiangcombineitembatchupdateasyncAPIRequest) GetCombineItemBatchUpsertAsyncRequest() *CombineItemBatchUpsertAsyncRequest {
	return r._combineItemBatchUpsertAsyncRequest
}
