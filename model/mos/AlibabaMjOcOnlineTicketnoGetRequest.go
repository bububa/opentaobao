package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线上小票号获取 API请求
alibaba.mj.oc.online.ticketno.get

线上小票号获取
*/
type AlibabaMjOcOnlineTicketnoGetRequest struct {
    model.Params
    // 外部门店号
    outStoreNo   string
}

// 初始化AlibabaMjOcOnlineTicketnoGetRequest对象
func NewAlibabaMjOcOnlineTicketnoGetRequest() *AlibabaMjOcOnlineTicketnoGetRequest{
    return &AlibabaMjOcOnlineTicketnoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMjOcOnlineTicketnoGetRequest) GetApiMethodName() string {
    return "alibaba.mj.oc.online.ticketno.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMjOcOnlineTicketnoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutStoreNo Setter
// 外部门店号
func (r *AlibabaMjOcOnlineTicketnoGetRequest) SetOutStoreNo(outStoreNo string) error {
    r.outStoreNo = outStoreNo
    r.Set("out_store_no", outStoreNo)
    return nil
}

// OutStoreNo Getter
func (r AlibabaMjOcOnlineTicketnoGetRequest) GetOutStoreNo() string {
    return r.outStoreNo
}
