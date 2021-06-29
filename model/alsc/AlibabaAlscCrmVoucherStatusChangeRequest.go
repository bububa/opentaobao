package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券状态更改 API请求
alibaba.alsc.crm.voucher.status.change

核销优惠券
*/
type AlibabaAlscCrmVoucherStatusChangeRequest struct {
    model.Params
    // 参数
    paramVoucherStatusChangeOpenReq   *VoucherStatusChangeOpenReq
}

// 初始化AlibabaAlscCrmVoucherStatusChangeRequest对象
func NewAlibabaAlscCrmVoucherStatusChangeRequest() *AlibabaAlscCrmVoucherStatusChangeRequest{
    return &AlibabaAlscCrmVoucherStatusChangeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmVoucherStatusChangeRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.voucher.status.change"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmVoucherStatusChangeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamVoucherStatusChangeOpenReq Setter
// 参数
func (r *AlibabaAlscCrmVoucherStatusChangeRequest) SetParamVoucherStatusChangeOpenReq(paramVoucherStatusChangeOpenReq *VoucherStatusChangeOpenReq) error {
    r.paramVoucherStatusChangeOpenReq = paramVoucherStatusChangeOpenReq
    r.Set("param_voucher_status_change_open_req", paramVoucherStatusChangeOpenReq)
    return nil
}

// ParamVoucherStatusChangeOpenReq Getter
func (r AlibabaAlscCrmVoucherStatusChangeRequest) GetParamVoucherStatusChangeOpenReq() *VoucherStatusChangeOpenReq {
    return r.paramVoucherStatusChangeOpenReq
}
