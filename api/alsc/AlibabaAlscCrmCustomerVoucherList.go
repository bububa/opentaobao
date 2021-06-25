package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
获取顾客优惠券列表 
alibaba.alsc.crm.customer.voucher.list

获取顾客优惠券列表
*/
func AlibabaAlscCrmCustomerVoucherList(clt *core.SDKClient, req *alsc.AlibabaAlscCrmCustomerVoucherListRequest, session string) (*alsc.AlibabaAlscCrmCustomerVoucherListResponse, error) {
    var resp alsc.AlibabaAlscCrmCustomerVoucherListAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
