package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店商品批量查询接口 API请求
alibaba.wdk.sku.storesku.scroll.query

提供门店商品批量查询接口
*/
type AlibabaWdkSkuStoreskuScrollQueryRequest struct {
    model.Params
    // 经营的id
    _storeId   string
    // 历游标，首次调用传递null，后续传递ScrollPageResult.getScrollId()
    _scrollId   string
}

// 初始化AlibabaWdkSkuStoreskuScrollQueryRequest对象
func NewAlibabaWdkSkuStoreskuScrollQueryRequest() *AlibabaWdkSkuStoreskuScrollQueryRequest{
    return &AlibabaWdkSkuStoreskuScrollQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuStoreskuScrollQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.storesku.scroll.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuStoreskuScrollQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 经营的id
func (r *AlibabaWdkSkuStoreskuScrollQueryRequest) SetStoreId(_storeId string) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r AlibabaWdkSkuStoreskuScrollQueryRequest) GetStoreId() string {
    return r._storeId
}
// ScrollId Setter
// 历游标，首次调用传递null，后续传递ScrollPageResult.getScrollId()
func (r *AlibabaWdkSkuStoreskuScrollQueryRequest) SetScrollId(_scrollId string) error {
    r._scrollId = _scrollId
    r.Set("scroll_id", _scrollId)
    return nil
}

// ScrollId Getter
func (r AlibabaWdkSkuStoreskuScrollQueryRequest) GetScrollId() string {
    return r._scrollId
}
