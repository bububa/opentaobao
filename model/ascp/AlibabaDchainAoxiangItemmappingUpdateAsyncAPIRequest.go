package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangitemmappingupdateasyncAPIRequest 创建/更新商货品关联关系 API请求
// alibaba.dchain.aoxiang.itemmapping.update.async
//
// 创建/更新商货品关联关系
type AlibabadchainaoxiangitemmappingupdateasyncAPIRequest struct {
	model.Params
	// 请求入参
	_itemMappingUpdateRequest *ItemMappingUpdateAsyncRequest
}

// NewAlibabadchainaoxiangitemmappingupdateasyncRequest 初始化AlibabadchainaoxiangitemmappingupdateasyncAPIRequest对象
func NewAlibabadchainaoxiangitemmappingupdateasyncRequest() *AlibabadchainaoxiangitemmappingupdateasyncAPIRequest {
	return &AlibabadchainaoxiangitemmappingupdateasyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangitemmappingupdateasyncAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.itemmapping.update.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangitemmappingupdateasyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangitemmappingupdateasyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemMappingUpdateRequest is ItemMappingUpdateRequest Setter
// 请求入参
func (r *AlibabadchainaoxiangitemmappingupdateasyncAPIRequest) SetItemMappingUpdateRequest(_itemMappingUpdateRequest *ItemMappingUpdateAsyncRequest) error {
	r._itemMappingUpdateRequest = _itemMappingUpdateRequest
	r.Set("item_mapping_update_request", _itemMappingUpdateRequest)
	return nil
}

// GetItemMappingUpdateRequest ItemMappingUpdateRequest Getter
func (r AlibabadchainaoxiangitemmappingupdateasyncAPIRequest) GetItemMappingUpdateRequest() *ItemMappingUpdateAsyncRequest {
	return r._itemMappingUpdateRequest
}
