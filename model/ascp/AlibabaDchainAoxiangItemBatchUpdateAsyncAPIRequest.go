package ascp

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangItemBatchUpdateAsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.item.batch.update.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangItemBatchUpdateAsyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
