package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店商品批量查询接口 APIRequest
alibaba.wdk.sku.storesku.scroll.query

提供门店商品批量查询接口
*/
type AlibabaWdkSkuStoreskuScrollQueryRequest struct {
    model.Params

    // 经营的id
    storeId   string 

    // 历游标，首次调用传递null，后续传递ScrollPageResult.getScrollId()
    scrollId   string 

}

func NewAlibabaWdkSkuStoreskuScrollQueryRequest() *AlibabaWdkSkuStoreskuScrollQueryRequest{
    return &AlibabaWdkSkuStoreskuScrollQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkSkuStoreskuScrollQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.storesku.scroll.query"
}

func (r AlibabaWdkSkuStoreskuScrollQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkSkuStoreskuScrollQueryRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r AlibabaWdkSkuStoreskuScrollQueryRequest) GetStoreId() string {
    return r.storeId
}

func (r *AlibabaWdkSkuStoreskuScrollQueryRequest) SetScrollId(scrollId string) error {
    r.scrollId = scrollId
    r.Set("scroll_id", scrollId)
    return nil
}

func (r AlibabaWdkSkuStoreskuScrollQueryRequest) GetScrollId() string {
    return r.scrollId
}

