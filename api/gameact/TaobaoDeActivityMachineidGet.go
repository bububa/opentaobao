package gameact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/gameact"
)

/* 
获取设备号 
taobao.de.activity.machineid.get

获取机器设备id
*/
func TaobaoDeActivityMachineidGet(clt *core.SDKClient, req *gameact.TaobaoDeActivityMachineidGetAPIRequest, session string) (*gameact.TaobaoDeActivityMachineidGetAPIResponse, error) {
    var resp gameact.TaobaoDeActivityMachineidGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
