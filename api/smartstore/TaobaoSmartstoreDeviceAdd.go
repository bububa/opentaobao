package smartstore

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/smartstore"
)

/* 
智慧门店设备创建 
taobao.smartstore.device.add

智慧门店设备创建
*/
func TaobaoSmartstoreDeviceAdd(clt *core.SDKClient, req *smartstore.TaobaoSmartstoreDeviceAddAPIRequest, session string) (*smartstore.TaobaoSmartstoreDeviceAddAPIResponse, error) {
    var resp smartstore.TaobaoSmartstoreDeviceAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
