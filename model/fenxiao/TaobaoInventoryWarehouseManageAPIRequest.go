package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoInventoryWarehouseManageAPIRequest 创建商家仓或者更新商家仓信息 API请求
// taobao.inventory.warehouse.manage
//
// 创建商家仓或者更新商家仓信息
type TaobaoInventoryWarehouseManageAPIRequest struct {
	model.Params
	// 仓库信息
	_wareHouseDto *WareHouseDto
}

// NewTaobaoInventoryWarehouseManageRequest 初始化TaobaoInventoryWarehouseManageAPIRequest对象
func NewTaobaoInventoryWarehouseManageRequest() *TaobaoInventoryWarehouseManageAPIRequest {
	return &TaobaoInventoryWarehouseManageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoInventoryWarehouseManageAPIRequest) GetApiMethodName() string {
	return "taobao.inventory.warehouse.manage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoInventoryWarehouseManageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoInventoryWarehouseManageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWareHouseDto is WareHouseDto Setter
// 仓库信息
func (r *TaobaoInventoryWarehouseManageAPIRequest) SetWareHouseDto(_wareHouseDto *WareHouseDto) error {
	r._wareHouseDto = _wareHouseDto
	r.Set("ware_house_dto", _wareHouseDto)
	return nil
}

// GetWareHouseDto WareHouseDto Getter
func (r TaobaoInventoryWarehouseManageAPIRequest) GetWareHouseDto() *WareHouseDto {
	return r._wareHouseDto
}
