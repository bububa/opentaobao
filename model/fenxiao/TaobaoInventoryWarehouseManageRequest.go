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
type TaobaoInventoryWarehouseManageRequest struct {
    model.Params
    // 仓库信息
    wareHouseDto   *WareHouseDto
}

// 初始化TaobaoInventoryWarehouseManageRequest对象
func NewTaobaoInventoryWarehouseManageRequest() *TaobaoInventoryWarehouseManageRequest{
    return &TaobaoInventoryWarehouseManageRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoInventoryWarehouseManageRequest) GetApiMethodName() string {
    return "taobao.inventory.warehouse.manage"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoInventoryWarehouseManageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WareHouseDto Setter
// 仓库信息
func (r *TaobaoInventoryWarehouseManageRequest) SetWareHouseDto(wareHouseDto *WareHouseDto) error {
    r.wareHouseDto = wareHouseDto
    r.Set("ware_house_dto", wareHouseDto)
    return nil
}

// WareHouseDto Getter
func (r TaobaoInventoryWarehouseManageRequest) GetWareHouseDto() *WareHouseDto {
    return r.wareHouseDto
}
