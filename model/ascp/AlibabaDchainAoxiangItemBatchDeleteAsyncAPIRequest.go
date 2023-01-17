package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangItemBatchDeleteAsyncAPIRequest 货品与组合货品删除 API请求
// alibaba.dchain.aoxiang.item.batch.delete.async
//
// 货品与组合货品删除
type AlibabaDchainAoxiangItemBatchDeleteAsyncAPIRequest struct {
	model.Params
	// 请求入参
	_itemDeleteRequest *ItemBatchDeleteAsyncRequest
}

// NewAlibabaDchainAoxiangItemBatchDeleteAsyncRequest 初始化AlibabaDchainAoxiangItemBatchDeleteAsyncAPIRequest对象
func NewAlibabaDchainAoxiangItemBatchDeleteAsyncRequest() *AlibabaDchainAoxiangItemBatchDeleteAsyncAPIRequest {
	return &AlibabaDchainAoxiangItemBatchDeleteAsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangItemBatchDeleteAsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.item.batch.delete.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangItemBatchDeleteAsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangItemBatchDeleteAsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemDeleteRequest is ItemDeleteRequest Setter
// 请求入参
func (r *AlibabaDchainAoxiangItemBatchDeleteAsyncAPIRequest) SetItemDeleteRequest(_itemDeleteRequest *ItemBatchDeleteAsyncRequest) error {
	r._itemDeleteRequest = _itemDeleteRequest
	r.Set("item_delete_request", _itemDeleteRequest)
	return nil
}

// GetItemDeleteRequest ItemDeleteRequest Getter
func (r AlibabaDchainAoxiangItemBatchDeleteAsyncAPIRequest) GetItemDeleteRequest() *ItemBatchDeleteAsyncRequest {
	return r._itemDeleteRequest
}
