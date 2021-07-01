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
type AlibabaAlscCrmVoucherStatusChangeAPIRequest struct {
    model.Params
    // 参数
    _paramVoucherStatusChangeOpenReq   *VoucherStatusChangeOpenReq
}

// 初始化AlibabaAlscCrmVoucherStatusChangeAPIRequest对象
func NewAlibabaAlscCrmVoucherStatusChangeRequest() *AlibabaAlscCrmVoucherStatusChangeAPIRequest{
    return &AlibabaAlscCrmVoucherStatusChangeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmVoucherStatusChangeAPIRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.voucher.status.change"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmVoucherStatusChangeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamVoucherStatusChangeOpenReq Setter
// 参数
func (r *AlibabaAlscCrmVoucherStatusChangeAPIRequest) SetParamVoucherStatusChangeOpenReq(_paramVoucherStatusChangeOpenReq *VoucherStatusChangeOpenReq) error {
    r._paramVoucherStatusChangeOpenReq = _paramVoucherStatusChangeOpenReq
    r.Set("param_voucher_status_change_open_req", _paramVoucherStatusChangeOpenReq)
    return nil
}

// ParamVoucherStatusChangeOpenReq Getter
func (r AlibabaAlscCrmVoucherStatusChangeAPIRequest) GetParamVoucherStatusChangeOpenReq() *VoucherStatusChangeOpenReq {
    return r._paramVoucherStatusChangeOpenReq
}
