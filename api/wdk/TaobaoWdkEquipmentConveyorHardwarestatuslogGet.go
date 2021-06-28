package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
硬件状态变化日志查询 
taobao.wdk.equipment.conveyor.hardwarestatuslog.get

硬件状态变化日志查询
*/
func TaobaoWdkEquipmentConveyorHardwarestatuslogGet(clt *core.SDKClient, req *wdk.TaobaoWdkEquipmentConveyorHardwarestatuslogGetRequest, session string) (*wdk.TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIResponse, error) {
    var resp wdk.TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
