package ottpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单查询接口(cp订单号查询) APIRequest
youku.ott.pay.order.queryorderbycp

给商户服务端查询订单状态
*/
type YoukuOttPayOrderQueryorderbycpRequest struct {
    model.Params

    // cp订单号
    cpOrderNo   string 

}

func NewYoukuOttPayOrderQueryorderbycpRequest() *YoukuOttPayOrderQueryorderbycpRequest{
    return &YoukuOttPayOrderQueryorderbycpRequest{
        Params: model.NewParams(),
    }
}

func (r YoukuOttPayOrderQueryorderbycpRequest) GetApiMethodName() string {
    return "youku.ott.pay.order.queryorderbycp"
}

func (r YoukuOttPayOrderQueryorderbycpRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YoukuOttPayOrderQueryorderbycpRequest) SetCpOrderNo(cpOrderNo string) error {
    r.cpOrderNo = cpOrderNo
    r.Set("cp_order_no", cpOrderNo)
    return nil
}

func (r YoukuOttPayOrderQueryorderbycpRequest) GetCpOrderNo() string {
    return r.cpOrderNo
}

