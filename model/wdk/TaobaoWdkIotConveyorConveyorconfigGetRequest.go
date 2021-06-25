package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取悬挂链基本配置信息 APIRequest
taobao.wdk.iot.conveyor.conveyorconfig.get

用于从云端WCS获取悬挂链基本配置信息
*/
type TaobaoWdkIotConveyorConveyorconfigGetRequest struct {
    model.Params

    // 仓编码
    warehouseCode   string 

    // 悬挂链id，默认为1
    conveyorId   int64 

}

func NewTaobaoWdkIotConveyorConveyorconfigGetRequest() *TaobaoWdkIotConveyorConveyorconfigGetRequest{
    return &TaobaoWdkIotConveyorConveyorconfigGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWdkIotConveyorConveyorconfigGetRequest) GetApiMethodName() string {
    return "taobao.wdk.iot.conveyor.conveyorconfig.get"
}

func (r TaobaoWdkIotConveyorConveyorconfigGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWdkIotConveyorConveyorconfigGetRequest) SetWarehouseCode(warehouseCode string) error {
    r.warehouseCode = warehouseCode
    r.Set("warehouse_code", warehouseCode)
    return nil
}

func (r TaobaoWdkIotConveyorConveyorconfigGetRequest) GetWarehouseCode() string {
    return r.warehouseCode
}

func (r *TaobaoWdkIotConveyorConveyorconfigGetRequest) SetConveyorId(conveyorId int64) error {
    r.conveyorId = conveyorId
    r.Set("conveyor_id", conveyorId)
    return nil
}

func (r TaobaoWdkIotConveyorConveyorconfigGetRequest) GetConveyorId() int64 {
    return r.conveyorId
}

