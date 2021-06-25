package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线上小票号获取 APIRequest
alibaba.mj.oc.online.ticketno.get

线上小票号获取
*/
type AlibabaMjOcOnlineTicketnoGetRequest struct {
    model.Params

    // 外部门店号
    outStoreNo   string 

}

func NewAlibabaMjOcOnlineTicketnoGetRequest() *AlibabaMjOcOnlineTicketnoGetRequest{
    return &AlibabaMjOcOnlineTicketnoGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMjOcOnlineTicketnoGetRequest) GetApiMethodName() string {
    return "alibaba.mj.oc.online.ticketno.get"
}

func (r AlibabaMjOcOnlineTicketnoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMjOcOnlineTicketnoGetRequest) SetOutStoreNo(outStoreNo string) error {
    r.outStoreNo = outStoreNo
    r.Set("out_store_no", outStoreNo)
    return nil
}

func (r AlibabaMjOcOnlineTicketnoGetRequest) GetOutStoreNo() string {
    return r.outStoreNo
}

