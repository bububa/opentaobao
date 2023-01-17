package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangItemmappingDeleteAPIRequest 删除商货品关联关系 API请求
// alibaba.dchain.aoxiang.itemmapping.delete
//
// 删除商货品关联关系
type AlibabaDchainAoxiangItemmappingDeleteAPIRequest struct {
	model.Params
	// 请求入参
	_itemMappingDeleteRequest *ItemMappingDeleteRequest
}

// NewAlibabaDchainAoxiangItemmappingDeleteRequest 初始化AlibabaDchainAoxiangItemmappingDeleteAPIRequest对象
func NewAlibabaDchainAoxiangItemmappingDeleteRequest() *AlibabaDchainAoxiangItemmappingDeleteAPIRequest {
	return &AlibabaDchainAoxiangItemmappingDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangItemmappingDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.itemmapping.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangItemmappingDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangItemmappingDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemMappingDeleteRequest is ItemMappingDeleteRequest Setter
// 请求入参
func (r *AlibabaDchainAoxiangItemmappingDeleteAPIRequest) SetItemMappingDeleteRequest(_itemMappingDeleteRequest *ItemMappingDeleteRequest) error {
	r._itemMappingDeleteRequest = _itemMappingDeleteRequest
	r.Set("item_mapping_delete_request", _itemMappingDeleteRequest)
	return nil
}

// GetItemMappingDeleteRequest ItemMappingDeleteRequest Getter
func (r AlibabaDchainAoxiangItemmappingDeleteAPIRequest) GetItemMappingDeleteRequest() *ItemMappingDeleteRequest {
	return r._itemMappingDeleteRequest
}
