package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发送券给指定用户 APIRequest
alibaba.alsc.crm.voucher.send

发送券给指定用户
*/
type AlibabaAlscCrmVoucherSendRequest struct {
    model.Params

    // 请求参数
    paramVoucherSendOpenReq   *VoucherSendOpenReq 

}

func NewAlibabaAlscCrmVoucherSendRequest() *AlibabaAlscCrmVoucherSendRequest{
    return &AlibabaAlscCrmVoucherSendRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmVoucherSendRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.voucher.send"
}

func (r AlibabaAlscCrmVoucherSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmVoucherSendRequest) SetParamVoucherSendOpenReq(paramVoucherSendOpenReq *VoucherSendOpenReq) error {
    r.paramVoucherSendOpenReq = paramVoucherSendOpenReq
    r.Set("param_voucher_send_open_req", paramVoucherSendOpenReq)
    return nil
}

func (r AlibabaAlscCrmVoucherSendRequest) GetParamVoucherSendOpenReq() *VoucherSendOpenReq {
    return r.paramVoucherSendOpenReq
}

