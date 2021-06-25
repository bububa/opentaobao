package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询汽车租赁活动信息 APIRequest
tmall.car.lease.item.activity.get

查询汽车租赁活动信息
*/
type TmallCarLeaseItemActivityGetRequest struct {
    model.Params

}

func NewTmallCarLeaseItemActivityGetRequest() *TmallCarLeaseItemActivityGetRequest{
    return &TmallCarLeaseItemActivityGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallCarLeaseItemActivityGetRequest) GetApiMethodName() string {
    return "tmall.car.lease.item.activity.get"
}

func (r TmallCarLeaseItemActivityGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


