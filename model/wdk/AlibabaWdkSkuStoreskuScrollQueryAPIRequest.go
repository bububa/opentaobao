package wdk

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuStoreskuScrollQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.storesku.scroll.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuStoreskuScrollQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StoreId Setter
// 经营的id
func (r *AlibabaWdkSkuStoreskuScrollQueryAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r AlibabaWdkSkuStoreskuScrollQueryAPIRequest) GetStoreId() string {
	return r._storeId
}

// Set is ScrollId Setter
// 历游标，首次调用传递null，后续传递ScrollPageResult.getScrollId()
func (r *AlibabaWdkSkuStoreskuScrollQueryAPIRequest) SetScrollId(_scrollId string) error {
	r._scrollId = _scrollId
	r.Set("scroll_id", _scrollId)
	return nil
}

// Get ScrollId Getter
func (r AlibabaWdkSkuStoreskuScrollQueryAPIRequest) GetScrollId() string {
	return r._scrollId
}
