package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取优惠券模版列表 API请求
alibaba.alsc.crm.voucher.template.list

获取优惠券模版列表
*/
type AlibabaAlscCrmVoucherTemplateListRequest struct {
    model.Params
    // 获取优惠模版规则请求参数
    _voucherTemplateOpenReq   *VoucherTemplateOpenReq
}

// 初始化AlibabaAlscCrmVoucherTemplateListRequest对象
func NewAlibabaAlscCrmVoucherTemplateListRequest() *AlibabaAlscCrmVoucherTemplateListRequest{
    return &AlibabaAlscCrmVoucherTemplateListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmVoucherTemplateListRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.voucher.template.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmVoucherTemplateListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// VoucherTemplateOpenReq Setter
// 获取优惠模版规则请求参数
func (r *AlibabaAlscCrmVoucherTemplateListRequest) SetVoucherTemplateOpenReq(_voucherTemplateOpenReq *VoucherTemplateOpenReq) error {
    r._voucherTemplateOpenReq = _voucherTemplateOpenReq
    r.Set("voucher_template_open_req", _voucherTemplateOpenReq)
    return nil
}

// VoucherTemplateOpenReq Getter
func (r AlibabaAlscCrmVoucherTemplateListRequest) GetVoucherTemplateOpenReq() *VoucherTemplateOpenReq {
    return r._voucherTemplateOpenReq
}
