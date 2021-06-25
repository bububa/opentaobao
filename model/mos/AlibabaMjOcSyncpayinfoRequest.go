package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
支付参考号回传 APIRequest
alibaba.mj.oc.syncpayinfo

支付参考号同步到oc
*/
type AlibabaMjOcSyncpayinfoRequest struct {
    model.Params

    // 支付参考号信息
    posPay   *PosPayDto 

}

func NewAlibabaMjOcSyncpayinfoRequest() *AlibabaMjOcSyncpayinfoRequest{
    return &AlibabaMjOcSyncpayinfoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMjOcSyncpayinfoRequest) GetApiMethodName() string {
    return "alibaba.mj.oc.syncpayinfo"
}

func (r AlibabaMjOcSyncpayinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMjOcSyncpayinfoRequest) SetPosPay(posPay *PosPayDto) error {
    r.posPay = posPay
    r.Set("pos_pay", posPay)
    return nil
}

func (r AlibabaMjOcSyncpayinfoRequest) GetPosPay() *PosPayDto {
    return r.posPay
}

