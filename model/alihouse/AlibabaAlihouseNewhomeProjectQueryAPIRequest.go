package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectQueryAPIRequest 查询楼盘相关信息 API请求
// alibaba.alihouse.newhome.project.query
//
// 根据outerid查询楼盘相关信息
type AlibabaAlihouseNewhomeProjectQueryAPIRequest struct {
	model.Params
	// 外部楼盘/小区id
	_outerId string
	// 商品id
	_itemId string
	// 门店ID
	_outerStoreId string
}

// NewAlibabaAlihouseNewhomeProjectQueryRequest 初始化AlibabaAlihouseNewhomeProjectQueryAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectQueryRequest() *AlibabaAlihouseNewhomeProjectQueryAPIRequest {
	return &AlibabaAlihouseNewhomeProjectQueryAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeProjectQueryAPIRequest) Reset() {
	r._outerId = ""
	r._itemId = ""
	r._outerStoreId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeProjectQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterId is OuterId Setter
// 外部楼盘/小区id
func (r *AlibabaAlihouseNewhomeProjectQueryAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaAlihouseNewhomeProjectQueryAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetItemId is ItemId Setter
// 商品id
func (r *AlibabaAlihouseNewhomeProjectQueryAPIRequest) SetItemId(_itemId string) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaAlihouseNewhomeProjectQueryAPIRequest) GetItemId() string {
	return r._itemId
}

// SetOuterStoreId is OuterStoreId Setter
// 门店ID
func (r *AlibabaAlihouseNewhomeProjectQueryAPIRequest) SetOuterStoreId(_outerStoreId string) error {
	r._outerStoreId = _outerStoreId
	r.Set("outer_store_id", _outerStoreId)
	return nil
}

// GetOuterStoreId OuterStoreId Getter
func (r AlibabaAlihouseNewhomeProjectQueryAPIRequest) GetOuterStoreId() string {
	return r._outerStoreId
}

var poolAlibabaAlihouseNewhomeProjectQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeProjectQueryRequest()
	},
}

// GetAlibabaAlihouseNewhomeProjectQueryRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectQueryAPIRequest
func GetAlibabaAlihouseNewhomeProjectQueryAPIRequest() *AlibabaAlihouseNewhomeProjectQueryAPIRequest {
	return poolAlibabaAlihouseNewhomeProjectQueryAPIRequest.Get().(*AlibabaAlihouseNewhomeProjectQueryAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeProjectQueryAPIRequest 将 AlibabaAlihouseNewhomeProjectQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectQueryAPIRequest(v *AlibabaAlihouseNewhomeProjectQueryAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectQueryAPIRequest.Put(v)
}
