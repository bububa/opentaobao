package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV查询门店 APIRequest
alibaba.lsy.crm.activity.store.getstorelist

ISV查询门店
*/
type AlibabaLsyCrmActivityStoreGetstorelistRequest struct {
    model.Params

    // 系统自动生成
    queryStoreReq   *NrtQueryStoreReq 

}

func NewAlibabaLsyCrmActivityStoreGetstorelistRequest() *AlibabaLsyCrmActivityStoreGetstorelistRequest{
    return &AlibabaLsyCrmActivityStoreGetstorelistRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLsyCrmActivityStoreGetstorelistRequest) GetApiMethodName() string {
    return "alibaba.lsy.crm.activity.store.getstorelist"
}

func (r AlibabaLsyCrmActivityStoreGetstorelistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLsyCrmActivityStoreGetstorelistRequest) SetQueryStoreReq(queryStoreReq *NrtQueryStoreReq) error {
    r.queryStoreReq = queryStoreReq
    r.Set("query_store_req", queryStoreReq)
    return nil
}

func (r AlibabaLsyCrmActivityStoreGetstorelistRequest) GetQueryStoreReq() *NrtQueryStoreReq {
    return r.queryStoreReq
}

