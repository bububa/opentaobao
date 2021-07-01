package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建商家仓或者更新商家仓信息 API请求
taobao.inventory.warehouse.manage

创建商家仓或者更新商家仓信息
*/
type TaobaoInventoryWarehouseManageAPIRequest struct {
    model.Params
    // 仓库信息
    _wareHouseDto   *WareHouseDto
}

// 初始化TaobaoInventoryWarehouseManageAPIRequest对象
func NewTaobaoInventoryWarehouseManageRequest() *TaobaoInventoryWarehouseManageAPIRequest{
    return &TaobaoInventoryWarehouseManageAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoInventoryWarehouseManageAPIRequest) GetApiMethodName() string {
    return "taobao.inventory.warehouse.manage"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoInventoryWarehouseManageAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WareHouseDto Setter
// 仓库信息
func (r *TaobaoInventoryWarehouseManageAPIRequest) SetWareHouseDto(_wareHouseDto *WareHouseDto) error {
    r._wareHouseDto = _wareHouseDto
    r.Set("ware_house_dto", _wareHouseDto)
    return nil
}

// WareHouseDto Getter
func (r TaobaoInventoryWarehouseManageAPIRequest) GetWareHouseDto() *WareHouseDto {
    return r._wareHouseDto
}
