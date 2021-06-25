package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
五道口按订单号批量查询供应商退款单 
alibaba.wdk.supplier.refund.get

五道口按订单号批量查询供应商退款单
*/
func AlibabaWdkSupplierRefundGet(clt *core.SDKClient, req *wdk.AlibabaWdkSupplierRefundGetRequest, session string) (*wdk.AlibabaWdkSupplierRefundGetResponse, error) {
    var resp wdk.AlibabaWdkSupplierRefundGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
