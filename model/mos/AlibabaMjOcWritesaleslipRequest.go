package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
开票占库 APIRequest
alibaba.mj.oc.writesaleslip

开票占库
*/
type AlibabaMjOcWritesaleslipRequest struct {
    model.Params

    // 开票占库入参
    posSaleOrder   *PosSaleOrderDto 

}

func NewAlibabaMjOcWritesaleslipRequest() *AlibabaMjOcWritesaleslipRequest{
    return &AlibabaMjOcWritesaleslipRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMjOcWritesaleslipRequest) GetApiMethodName() string {
    return "alibaba.mj.oc.writesaleslip"
}

func (r AlibabaMjOcWritesaleslipRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMjOcWritesaleslipRequest) SetPosSaleOrder(posSaleOrder *PosSaleOrderDto) error {
    r.posSaleOrder = posSaleOrder
    r.Set("pos_sale_order", posSaleOrder)
    return nil
}

func (r AlibabaMjOcWritesaleslipRequest) GetPosSaleOrder() *PosSaleOrderDto {
    return r.posSaleOrder
}

