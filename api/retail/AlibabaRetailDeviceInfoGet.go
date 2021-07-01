package retail

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/retail"
)

/* 
贩卖机设备信息获取 
alibaba.retail.device.info.get

贩卖机设备信息获取
*/
func AlibabaRetailDeviceInfoGet(clt *core.SDKClient, req *retail.AlibabaRetailDeviceInfoGetAPIRequest, session string) (*retail.AlibabaRetailDeviceInfoGetAPIResponse, error) {
    var resp retail.AlibabaRetailDeviceInfoGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
