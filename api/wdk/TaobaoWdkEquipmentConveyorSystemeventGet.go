package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
获取悬挂链系统事件 
taobao.wdk.equipment.conveyor.systemevent.get

五道口悬挂链系统事件查询
*/
func TaobaoWdkEquipmentConveyorSystemeventGet(clt *core.SDKClient, req *wdk.TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest, session string) (*wdk.TaobaoWdkEquipmentConveyorSystemeventGetAPIResponse, error) {
    var resp wdk.TaobaoWdkEquipmentConveyorSystemeventGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
