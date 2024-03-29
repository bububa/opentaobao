package ascp

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainAoxiangItemmappingUpdateAsyncAPIRequest) Reset() {
	r._itemMappingUpdateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangItemmappingUpdateAsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.itemmapping.update.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangItemmappingUpdateAsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangItemmappingUpdateAsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaDchainAoxiangItemmappingUpdateAsyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainAoxiangItemmappingUpdateAsyncRequest()
	},
}

// GetAlibabaDchainAoxiangItemmappingUpdateAsyncRequest 从 sync.Pool 获取 AlibabaDchainAoxiangItemmappingUpdateAsyncAPIRequest
func GetAlibabaDchainAoxiangItemmappingUpdateAsyncAPIRequest() *AlibabaDchainAoxiangItemmappingUpdateAsyncAPIRequest {
	return poolAlibabaDchainAoxiangItemmappingUpdateAsyncAPIRequest.Get().(*AlibabaDchainAoxiangItemmappingUpdateAsyncAPIRequest)
}

// ReleaseAlibabaDchainAoxiangItemmappingUpdateAsyncAPIRequest 将 AlibabaDchainAoxiangItemmappingUpdateAsyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainAoxiangItemmappingUpdateAsyncAPIRequest(v *AlibabaDchainAoxiangItemmappingUpdateAsyncAPIRequest) {
	v.Reset()
	poolAlibabaDchainAoxiangItemmappingUpdateAsyncAPIRequest.Put(v)
}
