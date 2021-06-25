package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据天猫id查询门店信息 APIRequest
alibaba.ssc.servicecenter.servicestore.query

根据天猫id查询门店信息
*/
type AlibabaSscServicecenterServicestoreQueryRequest struct {
    model.Params

    // 天猫id
    id   int64 

}

func NewAlibabaSscServicecenterServicestoreQueryRequest() *AlibabaSscServicecenterServicestoreQueryRequest{
    return &AlibabaSscServicecenterServicestoreQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSscServicecenterServicestoreQueryRequest) GetApiMethodName() string {
    return "alibaba.ssc.servicecenter.servicestore.query"
}

func (r AlibabaSscServicecenterServicestoreQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSscServicecenterServicestoreQueryRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r AlibabaSscServicecenterServicestoreQueryRequest) GetId() int64 {
    return r.id
}

