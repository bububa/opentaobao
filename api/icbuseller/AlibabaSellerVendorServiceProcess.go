package icbuseller

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbuseller"
)

/* 
服务商客户关联信息 
alibaba.seller.vendor.service.process

服务商客户关联信息
*/
func AlibabaSellerVendorServiceProcess(clt *core.SDKClient, req *icbuseller.AlibabaSellerVendorServiceProcessAPIRequest, session string) (*icbuseller.AlibabaSellerVendorServiceProcessAPIResponse, error) {
    var resp icbuseller.AlibabaSellerVendorServiceProcessAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
