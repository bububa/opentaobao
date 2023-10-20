package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoinventorywarehousemanageAPIRequest 创建商家仓或者更新商家仓信息 API请求
// taobao.inventory.warehouse.manage
//
// 创建商家仓或者更新商家仓信息
type TaobaoinventorywarehousemanageAPIRequest struct {
	model.Params
	// 仓库信息
	_wareHouseDto *WareHouseDto
}

// NewTaobaoinventorywarehousemanageRequest 初始化TaobaoinventorywarehousemanageAPIRequest对象
func NewTaobaoinventorywarehousemanageRequest() *TaobaoinventorywarehousemanageAPIRequest {
	return &TaobaoinventorywarehousemanageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoinventorywarehousemanageAPIRequest) GetApiMethodName() string {
	return "taobao.inventory.warehouse.manage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoinventorywarehousemanageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoinventorywarehousemanageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWareHouseDto is WareHouseDto Setter
// 仓库信息
func (r *TaobaoinventorywarehousemanageAPIRequest) SetWareHouseDto(_wareHouseDto *WareHouseDto) error {
	r._wareHouseDto = _wareHouseDto
	r.Set("ware_house_dto", _wareHouseDto)
	return nil
}

// GetWareHouseDto WareHouseDto Getter
func (r TaobaoinventorywarehousemanageAPIRequest) GetWareHouseDto() *WareHouseDto {
	return r._wareHouseDto
}
