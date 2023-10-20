package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangitembatchupdateasyncAPIRequest 货品新建/更新接口 API请求
// alibaba.dchain.aoxiang.item.batch.update.async
//
// 货品新建/更新接口
type AlibabadchainaoxiangitembatchupdateasyncAPIRequest struct {
	model.Params
	// 请求入参
	_itemUpdateRequest *ItemBatchUpdateAsyncRequest
}

// NewAlibabadchainaoxiangitembatchupdateasyncRequest 初始化AlibabadchainaoxiangitembatchupdateasyncAPIRequest对象
func NewAlibabadchainaoxiangitembatchupdateasyncRequest() *AlibabadchainaoxiangitembatchupdateasyncAPIRequest {
	return &AlibabadchainaoxiangitembatchupdateasyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangitembatchupdateasyncAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.item.batch.update.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangitembatchupdateasyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangitembatchupdateasyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemUpdateRequest is ItemUpdateRequest Setter
// 请求入参
func (r *AlibabadchainaoxiangitembatchupdateasyncAPIRequest) SetItemUpdateRequest(_itemUpdateRequest *ItemBatchUpdateAsyncRequest) error {
	r._itemUpdateRequest = _itemUpdateRequest
	r.Set("item_update_request", _itemUpdateRequest)
	return nil
}

// GetItemUpdateRequest ItemUpdateRequest Getter
func (r AlibabadchainaoxiangitembatchupdateasyncAPIRequest) GetItemUpdateRequest() *ItemBatchUpdateAsyncRequest {
	return r._itemUpdateRequest
}
