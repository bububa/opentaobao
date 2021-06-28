package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
阿里通信运营商信息回传 
alibaba.base.order.supplier.notify

接收阿里通信流量运营商信息回传
*/
func AlibabaBaseOrderSupplierNotify(clt *core.SDKClient, req *alicom.AlibabaBaseOrderSupplierNotifyRequest, session string) (*alicom.AlibabaBaseOrderSupplierNotifyAPIResponse, error) {
    var resp alicom.AlibabaBaseOrderSupplierNotifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
