package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangItemmappingUpdateAsyncAPIRequest 创建/更新商货品关联关系 API请求
// alibaba.dchain.aoxiang.itemmapping.update.async
//
// 创建/更新商货品关联关系
type AlibabaDchainAoxiangItemmappingUpdateAsyncAPIRequest struct {
	model.Params
	// 请求入参
	_itemMappingUpdateRequest *ItemMappingUpdateAsyncRequest
}

// NewAlibabaDchainAoxiangItemmappingUpdateAsyncRequest 初始化AlibabaDchainAoxiangItemmappingUpdateAsyncAPIRequest对象
func NewAlibabaDchainAoxiangItemmappingUpdateAsyncRequest() *AlibabaDchainAoxiangItemmappingUpdateAsyncAPIRequest {
	return &AlibabaDchainAoxiangItemmappingUpdateAsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangItemmappingUpdateAsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.itemmapping.update.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangItemmappingUpdateAsyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetItemMappingUpdateRequest is ItemMappingUpdateRequest Setter
// 请求入参
func (r *AlibabaDchainAoxiangItemmappingUpdateAsyncAPIRequest) SetItemMappingUpdateRequest(_itemMappingUpdateRequest *ItemMappingUpdateAsyncRequest) error {
	r._itemMappingUpdateRequest = _itemMappingUpdateRequest
	r.Set("item_mapping_update_request", _itemMappingUpdateRequest)
	return nil
}

// GetItemMappingUpdateRequest ItemMappingUpdateRequest Getter
func (r AlibabaDchainAoxiangItemmappingUpdateAsyncAPIRequest) GetItemMappingUpdateRequest() *ItemMappingUpdateAsyncRequest {
	return r._itemMappingUpdateRequest
}
