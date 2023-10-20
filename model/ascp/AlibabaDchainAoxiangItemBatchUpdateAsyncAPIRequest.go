package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangItemBatchUpdateAsyncAPIRequest 货品新建/更新接口 API请求
// alibaba.dchain.aoxiang.item.batch.update.async
//
// 货品新建/更新接口
type AlibabaDchainAoxiangItemBatchUpdateAsyncAPIRequest struct {
	model.Params
	// 请求入参
	_itemUpdateRequest *ItemBatchUpdateAsyncRequest
}

// NewAlibabaDchainAoxiangItemBatchUpdateAsyncRequest 初始化AlibabaDchainAoxiangItemBatchUpdateAsyncAPIRequest对象
func NewAlibabaDchainAoxiangItemBatchUpdateAsyncRequest() *AlibabaDchainAoxiangItemBatchUpdateAsyncAPIRequest {
	return &AlibabaDchainAoxiangItemBatchUpdateAsyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainAoxiangItemBatchUpdateAsyncAPIRequest) Reset() {
	r._itemUpdateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangItemBatchUpdateAsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.item.batch.update.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangItemBatchUpdateAsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangItemBatchUpdateAsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemUpdateRequest is ItemUpdateRequest Setter
// 请求入参
func (r *AlibabaDchainAoxiangItemBatchUpdateAsyncAPIRequest) SetItemUpdateRequest(_itemUpdateRequest *ItemBatchUpdateAsyncRequest) error {
	r._itemUpdateRequest = _itemUpdateRequest
	r.Set("item_update_request", _itemUpdateRequest)
	return nil
}

// GetItemUpdateRequest ItemUpdateRequest Getter
func (r AlibabaDchainAoxiangItemBatchUpdateAsyncAPIRequest) GetItemUpdateRequest() *ItemBatchUpdateAsyncRequest {
	return r._itemUpdateRequest
}

var poolAlibabaDchainAoxiangItemBatchUpdateAsyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainAoxiangItemBatchUpdateAsyncRequest()
	},
}

// GetAlibabaDchainAoxiangItemBatchUpdateAsyncRequest 从 sync.Pool 获取 AlibabaDchainAoxiangItemBatchUpdateAsyncAPIRequest
func GetAlibabaDchainAoxiangItemBatchUpdateAsyncAPIRequest() *AlibabaDchainAoxiangItemBatchUpdateAsyncAPIRequest {
	return poolAlibabaDchainAoxiangItemBatchUpdateAsyncAPIRequest.Get().(*AlibabaDchainAoxiangItemBatchUpdateAsyncAPIRequest)
}

// ReleaseAlibabaDchainAoxiangItemBatchUpdateAsyncAPIRequest 将 AlibabaDchainAoxiangItemBatchUpdateAsyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainAoxiangItemBatchUpdateAsyncAPIRequest(v *AlibabaDchainAoxiangItemBatchUpdateAsyncAPIRequest) {
	v.Reset()
	poolAlibabaDchainAoxiangItemBatchUpdateAsyncAPIRequest.Put(v)
}
