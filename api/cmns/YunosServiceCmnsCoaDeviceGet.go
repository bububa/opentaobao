package cmns

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/cmns"
)

/* 
设备详情查询 
yunos.service.cmns.coa.device.get

第三方应用开发者调用此接口查询设备详情
*/
func YunosServiceCmnsCoaDeviceGet(clt *core.SDKClient, req *cmns.YunosServiceCmnsCoaDeviceGetRequest, session string) (*cmns.YunosServiceCmnsCoaDeviceGetResponse, error) {
    var resp cmns.YunosServiceCmnsCoaDeviceGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
