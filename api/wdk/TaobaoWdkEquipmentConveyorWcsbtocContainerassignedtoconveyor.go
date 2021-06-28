package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
容器被预分拣器分配到悬挂链 
taobao.wdk.equipment.conveyor.wcsbtoc.containerassignedtoconveyor

容器被预分拣器分配到悬挂链
*/
func TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyor(clt *core.SDKClient, req *wdk.TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorRequest, session string) (*wdk.TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIResponse, error) {
    var resp wdk.TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
