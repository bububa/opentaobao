package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券状态更改 APIRequest
alibaba.alsc.crm.voucher.status.change

核销优惠券
*/
type AlibabaAlscCrmVoucherStatusChangeRequest struct {
    model.Params

    // 参数
    paramVoucherStatusChangeOpenReq   *VoucherStatusChangeOpenReq 

}

func NewAlibabaAlscCrmVoucherStatusChangeRequest() *AlibabaAlscCrmVoucherStatusChangeRequest{
    return &AlibabaAlscCrmVoucherStatusChangeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmVoucherStatusChangeRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.voucher.status.change"
}

func (r AlibabaAlscCrmVoucherStatusChangeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmVoucherStatusChangeRequest) SetParamVoucherStatusChangeOpenReq(paramVoucherStatusChangeOpenReq *VoucherStatusChangeOpenReq) error {
    r.paramVoucherStatusChangeOpenReq = paramVoucherStatusChangeOpenReq
    r.Set("param_voucher_status_change_open_req", paramVoucherStatusChangeOpenReq)
    return nil
}

func (r AlibabaAlscCrmVoucherStatusChangeRequest) GetParamVoucherStatusChangeOpenReq() *VoucherStatusChangeOpenReq {
    return r.paramVoucherStatusChangeOpenReq
}

