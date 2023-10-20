package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuStoreskuScrollQueryAPIRequest 门店商品批量查询接口 API请求
// alibaba.wdk.sku.storesku.scroll.query
//
// 提供门店商品批量查询接口
type AlibabaWdkSkuStoreskuScrollQueryAPIRequest struct {
	model.Params
	// 经营的id
	_storeId string
	// 历游标，首次调用传递null，后续传递ScrollPageResult.getScrollId()
	_scrollId string
}

// NewAlibabaWdkSkuStoreskuScrollQueryRequest 初始化AlibabaWdkSkuStoreskuScrollQueryAPIRequest对象
func NewAlibabaWdkSkuStoreskuScrollQueryRequest() *AlibabaWdkSkuStoreskuScrollQueryAPIRequest {
	return &AlibabaWdkSkuStoreskuScrollQueryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkSkuStoreskuScrollQueryAPIRequest) Reset() {
	r._storeId = ""
	r._scrollId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuStoreskuScrollQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.storesku.scroll.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuStoreskuScrollQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkSkuStoreskuScrollQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 经营的id
func (r *AlibabaWdkSkuStoreskuScrollQueryAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabaWdkSkuStoreskuScrollQueryAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetScrollId is ScrollId Setter
// 历游标，首次调用传递null，后续传递ScrollPageResult.getScrollId()
func (r *AlibabaWdkSkuStoreskuScrollQueryAPIRequest) SetScrollId(_scrollId string) error {
	r._scrollId = _scrollId
	r.Set("scroll_id", _scrollId)
	return nil
}

// GetScrollId ScrollId Getter
func (r AlibabaWdkSkuStoreskuScrollQueryAPIRequest) GetScrollId() string {
	return r._scrollId
}

var poolAlibabaWdkSkuStoreskuScrollQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkSkuStoreskuScrollQueryRequest()
	},
}

// GetAlibabaWdkSkuStoreskuScrollQueryRequest 从 sync.Pool 获取 AlibabaWdkSkuStoreskuScrollQueryAPIRequest
func GetAlibabaWdkSkuStoreskuScrollQueryAPIRequest() *AlibabaWdkSkuStoreskuScrollQueryAPIRequest {
	return poolAlibabaWdkSkuStoreskuScrollQueryAPIRequest.Get().(*AlibabaWdkSkuStoreskuScrollQueryAPIRequest)
}

// ReleaseAlibabaWdkSkuStoreskuScrollQueryAPIRequest 将 AlibabaWdkSkuStoreskuScrollQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkSkuStoreskuScrollQueryAPIRequest(v *AlibabaWdkSkuStoreskuScrollQueryAPIRequest) {
	v.Reset()
	poolAlibabaWdkSkuStoreskuScrollQueryAPIRequest.Put(v)
}
