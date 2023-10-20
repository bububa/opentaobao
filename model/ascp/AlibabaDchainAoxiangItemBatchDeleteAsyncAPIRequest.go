package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangitembatchdeleteasyncAPIRequest 货品与组合货品删除 API请求
// alibaba.dchain.aoxiang.item.batch.delete.async
//
// 货品与组合货品删除
type AlibabadchainaoxiangitembatchdeleteasyncAPIRequest struct {
	model.Params
	// 请求入参
	_itemDeleteRequest *ItemBatchDeleteAsyncRequest
}

// NewAlibabadchainaoxiangitembatchdeleteasyncRequest 初始化AlibabadchainaoxiangitembatchdeleteasyncAPIRequest对象
func NewAlibabadchainaoxiangitembatchdeleteasyncRequest() *AlibabadchainaoxiangitembatchdeleteasyncAPIRequest {
	return &AlibabadchainaoxiangitembatchdeleteasyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangitembatchdeleteasyncAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.item.batch.delete.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangitembatchdeleteasyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangitembatchdeleteasyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemDeleteRequest is ItemDeleteRequest Setter
// 请求入参
func (r *AlibabadchainaoxiangitembatchdeleteasyncAPIRequest) SetItemDeleteRequest(_itemDeleteRequest *ItemBatchDeleteAsyncRequest) error {
	r._itemDeleteRequest = _itemDeleteRequest
	r.Set("item_delete_request", _itemDeleteRequest)
	return nil
}

// GetItemDeleteRequest ItemDeleteRequest Getter
func (r AlibabadchainaoxiangitembatchdeleteasyncAPIRequest) GetItemDeleteRequest() *ItemBatchDeleteAsyncRequest {
	return r._itemDeleteRequest
}
