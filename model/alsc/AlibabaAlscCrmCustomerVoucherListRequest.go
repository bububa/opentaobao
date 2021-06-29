package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取顾客优惠券列表 API请求
alibaba.alsc.crm.customer.voucher.list

获取顾客优惠券列表
*/
type AlibabaAlscCrmCustomerVoucherListRequest struct {
    model.Params
    // 获取顾客优惠券列表
    customerVoucherFullOpenReq   *CustomerVoucherFullOpenReq
}

// 初始化AlibabaAlscCrmCustomerVoucherListRequest对象
func NewAlibabaAlscCrmCustomerVoucherListRequest() *AlibabaAlscCrmCustomerVoucherListRequest{
    return &AlibabaAlscCrmCustomerVoucherListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCustomerVoucherListRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.customer.voucher.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCustomerVoucherListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CustomerVoucherFullOpenReq Setter
// 获取顾客优惠券列表
func (r *AlibabaAlscCrmCustomerVoucherListRequest) SetCustomerVoucherFullOpenReq(customerVoucherFullOpenReq *CustomerVoucherFullOpenReq) error {
    r.customerVoucherFullOpenReq = customerVoucherFullOpenReq
    r.Set("customer_voucher_full_open_req", customerVoucherFullOpenReq)
    return nil
}

// CustomerVoucherFullOpenReq Getter
func (r AlibabaAlscCrmCustomerVoucherListRequest) GetCustomerVoucherFullOpenReq() *CustomerVoucherFullOpenReq {
    return r.customerVoucherFullOpenReq
}
