package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建商家仓或者更新商家仓信息 APIRequest
taobao.inventory.warehouse.manage

创建商家仓或者更新商家仓信息
*/
type TaobaoInventoryWarehouseManageRequest struct {
    model.Params

    // 仓库信息
    wareHouseDto   *WareHouseDto 

}

func NewTaobaoInventoryWarehouseManageRequest() *TaobaoInventoryWarehouseManageRequest{
    return &TaobaoInventoryWarehouseManageRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoInventoryWarehouseManageRequest) GetApiMethodName() string {
    return "taobao.inventory.warehouse.manage"
}

func (r TaobaoInventoryWarehouseManageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoInventoryWarehouseManageRequest) SetWareHouseDto(wareHouseDto *WareHouseDto) error {
    r.wareHouseDto = wareHouseDto
    r.Set("ware_house_dto", wareHouseDto)
    return nil
}

func (r TaobaoInventoryWarehouseManageRequest) GetWareHouseDto() *WareHouseDto {
    return r.wareHouseDto
}

