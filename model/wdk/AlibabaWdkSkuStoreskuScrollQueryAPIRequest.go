package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkskustoreskuscrollqueryAPIRequest 门店商品批量查询接口 API请求
// alibaba.wdk.sku.storesku.scroll.query
//
// 提供门店商品批量查询接口
type AlibabawdkskustoreskuscrollqueryAPIRequest struct {
	model.Params
	// 经营的id
	_storeId string
	// 历游标，首次调用传递null，后续传递ScrollPageResult.getScrollId()
	_scrollId string
}

// NewAlibabawdkskustoreskuscrollqueryRequest 初始化AlibabawdkskustoreskuscrollqueryAPIRequest对象
func NewAlibabawdkskustoreskuscrollqueryRequest() *AlibabawdkskustoreskuscrollqueryAPIRequest {
	return &AlibabawdkskustoreskuscrollqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkskustoreskuscrollqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.storesku.scroll.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkskustoreskuscrollqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkskustoreskuscrollqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 经营的id
func (r *AlibabawdkskustoreskuscrollqueryAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabawdkskustoreskuscrollqueryAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetScrollId is ScrollId Setter
// 历游标，首次调用传递null，后续传递ScrollPageResult.getScrollId()
func (r *AlibabawdkskustoreskuscrollqueryAPIRequest) SetScrollId(_scrollId string) error {
	r._scrollId = _scrollId
	r.Set("scroll_id", _scrollId)
	return nil
}

// GetScrollId ScrollId Getter
func (r AlibabawdkskustoreskuscrollqueryAPIRequest) GetScrollId() string {
	return r._scrollId
}
