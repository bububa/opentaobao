package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取顾客优惠券列表 APIRequest
alibaba.alsc.crm.customer.voucher.list

获取顾客优惠券列表
*/
type AlibabaAlscCrmCustomerVoucherListRequest struct {
    model.Params

    // 获取顾客优惠券列表
    customerVoucherFullOpenReq   *CustomerVoucherFullOpenReq 

}

func NewAlibabaAlscCrmCustomerVoucherListRequest() *AlibabaAlscCrmCustomerVoucherListRequest{
    return &AlibabaAlscCrmCustomerVoucherListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmCustomerVoucherListRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.customer.voucher.list"
}

func (r AlibabaAlscCrmCustomerVoucherListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmCustomerVoucherListRequest) SetCustomerVoucherFullOpenReq(customerVoucherFullOpenReq *CustomerVoucherFullOpenReq) error {
    r.customerVoucherFullOpenReq = customerVoucherFullOpenReq
    r.Set("customer_voucher_full_open_req", customerVoucherFullOpenReq)
    return nil
}

func (r AlibabaAlscCrmCustomerVoucherListRequest) GetCustomerVoucherFullOpenReq() *CustomerVoucherFullOpenReq {
    return r.customerVoucherFullOpenReq
}

