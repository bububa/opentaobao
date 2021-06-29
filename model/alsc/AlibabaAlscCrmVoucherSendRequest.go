package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发送券给指定用户 API请求
alibaba.alsc.crm.voucher.send

发送券给指定用户
*/
type AlibabaAlscCrmVoucherSendRequest struct {
    model.Params
    // 请求参数
    _paramVoucherSendOpenReq   *VoucherSendOpenReq
}

// 初始化AlibabaAlscCrmVoucherSendRequest对象
func NewAlibabaAlscCrmVoucherSendRequest() *AlibabaAlscCrmVoucherSendRequest{
    return &AlibabaAlscCrmVoucherSendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmVoucherSendRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.voucher.send"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmVoucherSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamVoucherSendOpenReq Setter
// 请求参数
func (r *AlibabaAlscCrmVoucherSendRequest) SetParamVoucherSendOpenReq(_paramVoucherSendOpenReq *VoucherSendOpenReq) error {
    r._paramVoucherSendOpenReq = _paramVoucherSendOpenReq
    r.Set("param_voucher_send_open_req", _paramVoucherSendOpenReq)
    return nil
}

// ParamVoucherSendOpenReq Getter
func (r AlibabaAlscCrmVoucherSendRequest) GetParamVoucherSendOpenReq() *VoucherSendOpenReq {
    return r._paramVoucherSendOpenReq
}
