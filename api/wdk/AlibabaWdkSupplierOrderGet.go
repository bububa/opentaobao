package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
五道口按订单号批量查询供应商正向订单 
alibaba.wdk.supplier.order.get

五道口按订单号批量查询供应商正向订单
*/
func AlibabaWdkSupplierOrderGet(clt *core.SDKClient, req *wdk.AlibabaWdkSupplierOrderGetAPIRequest, session string) (*wdk.AlibabaWdkSupplierOrderGetAPIResponse, error) {
    var resp wdk.AlibabaWdkSupplierOrderGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
