package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
悬挂链状态回传确认 
taobao.wdk.equipment.conveyor.statusconfirm

悬挂链状态回传确认
*/
func TaobaoWdkEquipmentConveyorStatusconfirm(clt *core.SDKClient, req *wdk.TaobaoWdkEquipmentConveyorStatusconfirmRequest, session string) (*wdk.TaobaoWdkEquipmentConveyorStatusconfirmResponse, error) {
    var resp wdk.TaobaoWdkEquipmentConveyorStatusconfirmAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
