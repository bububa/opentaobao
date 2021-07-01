package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取悬挂链基本配置信息 API请求
taobao.wdk.iot.conveyor.conveyorconfig.get

用于从云端WCS获取悬挂链基本配置信息
*/
type TaobaoWdkIotConveyorConveyorconfigGetAPIRequest struct {
    model.Params
    // 仓编码
    _warehouseCode   string
    // 悬挂链id，默认为1
    _conveyorId   int64
}

// 初始化TaobaoWdkIotConveyorConveyorconfigGetAPIRequest对象
func NewTaobaoWdkIotConveyorConveyorconfigGetRequest() *TaobaoWdkIotConveyorConveyorconfigGetAPIRequest{
    return &TaobaoWdkIotConveyorConveyorconfigGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWdkIotConveyorConveyorconfigGetAPIRequest) GetApiMethodName() string {
    return "taobao.wdk.iot.conveyor.conveyorconfig.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWdkIotConveyorConveyorconfigGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WarehouseCode Setter
// 仓编码
func (r *TaobaoWdkIotConveyorConveyorconfigGetAPIRequest) SetWarehouseCode(_warehouseCode string) error {
    r._warehouseCode = _warehouseCode
    r.Set("warehouse_code", _warehouseCode)
    return nil
}

// WarehouseCode Getter
func (r TaobaoWdkIotConveyorConveyorconfigGetAPIRequest) GetWarehouseCode() string {
    return r._warehouseCode
}
// ConveyorId Setter
// 悬挂链id，默认为1
func (r *TaobaoWdkIotConveyorConveyorconfigGetAPIRequest) SetConveyorId(_conveyorId int64) error {
    r._conveyorId = _conveyorId
    r.Set("conveyor_id", _conveyorId)
    return nil
}

// ConveyorId Getter
func (r TaobaoWdkIotConveyorConveyorconfigGetAPIRequest) GetConveyorId() int64 {
    return r._conveyorId
}
