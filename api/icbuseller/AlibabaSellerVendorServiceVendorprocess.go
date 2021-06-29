package icbuseller

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbuseller"
)

/* 
服务商客户关联信息 
alibaba.seller.vendor.service.vendorprocess

服务商客户关联信息
*/
func AlibabaSellerVendorServiceVendorprocess(clt *core.SDKClient, req *icbuseller.AlibabaSellerVendorServiceVendorprocessRequest, session string) (*icbuseller.AlibabaSellerVendorServiceVendorprocessAPIResponse, error) {
    var resp icbuseller.AlibabaSellerVendorServiceVendorprocessAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
