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
type AlibabaMjOcOnlineTicketnoGetAPIRequest struct {
    model.Params
    // 外部门店号
    _outStoreNo   string
}

// 初始化AlibabaMjOcOnlineTicketnoGetAPIRequest对象
func NewAlibabaMjOcOnlineTicketnoGetRequest() *AlibabaMjOcOnlineTicketnoGetAPIRequest{
    return &AlibabaMjOcOnlineTicketnoGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMjOcOnlineTicketnoGetAPIRequest) GetApiMethodName() string {
    return "alibaba.mj.oc.online.ticketno.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMjOcOnlineTicketnoGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutStoreNo Setter
// 外部门店号
func (r *AlibabaMjOcOnlineTicketnoGetAPIRequest) SetOutStoreNo(_outStoreNo string) error {
    r._outStoreNo = _outStoreNo
    r.Set("out_store_no", _outStoreNo)
    return nil
}

// OutStoreNo Getter
func (r AlibabaMjOcOnlineTicketnoGetAPIRequest) GetOutStoreNo() string {
    return r._outStoreNo
}
