package tmallcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫开新车查询订单id APIRequest
tmall.car.lease.orderid.get

天猫开新车查询订单id
*/
type TmallCarLeaseOrderidGetRequest struct {
    model.Params

    // openid
    openId   string 

}

func NewTmallCarLeaseOrderidGetRequest() *TmallCarLeaseOrderidGetRequest{
    return &TmallCarLeaseOrderidGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallCarLeaseOrderidGetRequest) GetApiMethodName() string {
    return "tmall.car.lease.orderid.get"
}

func (r TmallCarLeaseOrderidGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallCarLeaseOrderidGetRequest) SetOpenId(openId string) error {
    r.openId = openId
    r.Set("open_id", openId)
    return nil
}

func (r TmallCarLeaseOrderidGetRequest) GetOpenId() string {
    return r.openId
}

