package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
支付参考号回传 API请求
alibaba.mj.oc.syncpayinfo

支付参考号同步到oc
*/
type AlibabaMjOcSyncpayinfoRequest struct {
    model.Params
    // 支付参考号信息
    posPay   *PosPayDto
}

// 初始化AlibabaMjOcSyncpayinfoRequest对象
func NewAlibabaMjOcSyncpayinfoRequest() *AlibabaMjOcSyncpayinfoRequest{
    return &AlibabaMjOcSyncpayinfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMjOcSyncpayinfoRequest) GetApiMethodName() string {
    return "alibaba.mj.oc.syncpayinfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMjOcSyncpayinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PosPay Setter
// 支付参考号信息
func (r *AlibabaMjOcSyncpayinfoRequest) SetPosPay(posPay *PosPayDto) error {
    r.posPay = posPay
    r.Set("pos_pay", posPay)
    return nil
}

// PosPay Getter
func (r AlibabaMjOcSyncpayinfoRequest) GetPosPay() *PosPayDto {
    return r.posPay
}
