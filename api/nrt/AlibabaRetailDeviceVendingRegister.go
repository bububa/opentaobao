package nrt

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/nrt"
)

/* 
贩卖机设备注册 
alibaba.retail.device.vending.register

贩卖机注册
*/
func AlibabaRetailDeviceVendingRegister(clt *core.SDKClient, req *nrt.AlibabaRetailDeviceVendingRegisterAPIRequest, session string) (*nrt.AlibabaRetailDeviceVendingRegisterAPIResponse, error) {
    var resp nrt.AlibabaRetailDeviceVendingRegisterAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
