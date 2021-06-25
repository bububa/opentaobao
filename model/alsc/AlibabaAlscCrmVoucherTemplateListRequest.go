package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取优惠券模版列表 APIRequest
alibaba.alsc.crm.voucher.template.list

获取优惠券模版列表
*/
type AlibabaAlscCrmVoucherTemplateListRequest struct {
    model.Params

    // 获取优惠模版规则请求参数
    voucherTemplateOpenReq   *VoucherTemplateOpenReq 

}

func NewAlibabaAlscCrmVoucherTemplateListRequest() *AlibabaAlscCrmVoucherTemplateListRequest{
    return &AlibabaAlscCrmVoucherTemplateListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmVoucherTemplateListRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.voucher.template.list"
}

func (r AlibabaAlscCrmVoucherTemplateListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmVoucherTemplateListRequest) SetVoucherTemplateOpenReq(voucherTemplateOpenReq *VoucherTemplateOpenReq) error {
    r.voucherTemplateOpenReq = voucherTemplateOpenReq
    r.Set("voucher_template_open_req", voucherTemplateOpenReq)
    return nil
}

func (r AlibabaAlscCrmVoucherTemplateListRequest) GetVoucherTemplateOpenReq() *VoucherTemplateOpenReq {
    return r.voucherTemplateOpenReq
}

