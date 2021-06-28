package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
获取批次或波次中容器的信息 
taobao.wdk.equipment.conveyor.containerinfo.get

获取批次或波次中容器的信息
*/
func TaobaoWdkEquipmentConveyorContainerinfoGet(clt *core.SDKClient, req *wdk.TaobaoWdkEquipmentConveyorContainerinfoGetRequest, session string) (*wdk.TaobaoWdkEquipmentConveyorContainerinfoGetAPIResponse, error) {
    var resp wdk.TaobaoWdkEquipmentConveyorContainerinfoGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
