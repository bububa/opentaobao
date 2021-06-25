package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
获取五道口悬挂链信息 
taobao.wdk.equipment.conveyor.conveyorinfo.get

获取五道口悬挂链信息
*/
func TaobaoWdkEquipmentConveyorConveyorinfoGet(clt *core.SDKClient, req *wdk.TaobaoWdkEquipmentConveyorConveyorinfoGetRequest, session string) (*wdk.TaobaoWdkEquipmentConveyorConveyorinfoGetResponse, error) {
    var resp wdk.TaobaoWdkEquipmentConveyorConveyorinfoGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
