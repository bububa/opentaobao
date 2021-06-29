package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV查询门店 API请求
alibaba.lsy.crm.activity.store.getstorelist

ISV查询门店
*/
type AlibabaLsyCrmActivityStoreGetstorelistRequest struct {
    model.Params
    // 系统自动生成
    _queryStoreReq   *NrtQueryStoreReq
}

// 初始化AlibabaLsyCrmActivityStoreGetstorelistRequest对象
func NewAlibabaLsyCrmActivityStoreGetstorelistRequest() *AlibabaLsyCrmActivityStoreGetstorelistRequest{
    return &AlibabaLsyCrmActivityStoreGetstorelistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLsyCrmActivityStoreGetstorelistRequest) GetApiMethodName() string {
    return "alibaba.lsy.crm.activity.store.getstorelist"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLsyCrmActivityStoreGetstorelistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QueryStoreReq Setter
// 系统自动生成
func (r *AlibabaLsyCrmActivityStoreGetstorelistRequest) SetQueryStoreReq(_queryStoreReq *NrtQueryStoreReq) error {
    r._queryStoreReq = _queryStoreReq
    r.Set("query_store_req", _queryStoreReq)
    return nil
}

// QueryStoreReq Getter
func (r AlibabaLsyCrmActivityStoreGetstorelistRequest) GetQueryStoreReq() *NrtQueryStoreReq {
    return r._queryStoreReq
}
