package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
五道口悬挂链信息批量确认 
taobao.wdk.equipment.conveyor.batchconfirm

批量消息确认
*/
func TaobaoWdkEquipmentConveyorBatchconfirm(clt *core.SDKClient, req *wdk.TaobaoWdkEquipmentConveyorBatchconfirmRequest, session string) (*wdk.TaobaoWdkEquipmentConveyorBatchconfirmResponse, error) {
    var resp wdk.TaobaoWdkEquipmentConveyorBatchconfirmAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
