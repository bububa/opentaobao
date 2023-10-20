package lstpos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstposopeninventorysyncinventorydataAPIRequest 商品库存修改同步接口(最多20条库存信息) API请求
// alibaba.lst.pos.open.inventory.syncinventorydata
//
// 商品库存修改同步接口(最多20条库存信息)
type AlibabalstposopeninventorysyncinventorydataAPIRequest struct {
	model.Params
	// 库存对象列表
	_inventoryDTOList []InventoryDto
	// 门店对应的主账号id
	_userId int64
}

// NewAlibabalstposopeninventorysyncinventorydataRequest 初始化AlibabalstposopeninventorysyncinventorydataAPIRequest对象
func NewAlibabalstposopeninventorysyncinventorydataRequest() *AlibabalstposopeninventorysyncinventorydataAPIRequest {
	return &AlibabalstposopeninventorysyncinventorydataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalstposopeninventorysyncinventorydataAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.pos.open.inventory.syncinventorydata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalstposopeninventorysyncinventorydataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalstposopeninventorysyncinventorydataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInventoryDTOList is InventoryDTOList Setter
// 库存对象列表
func (r *AlibabalstposopeninventorysyncinventorydataAPIRequest) SetInventoryDTOList(_inventoryDTOList []InventoryDto) error {
	r._inventoryDTOList = _inventoryDTOList
	r.Set("inventory_d_t_o_list", _inventoryDTOList)
	return nil
}

// GetInventoryDTOList InventoryDTOList Getter
func (r AlibabalstposopeninventorysyncinventorydataAPIRequest) GetInventoryDTOList() []InventoryDto {
	return r._inventoryDTOList
}

// SetUserId is UserId Setter
// 门店对应的主账号id
func (r *AlibabalstposopeninventorysyncinventorydataAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabalstposopeninventorysyncinventorydataAPIRequest) GetUserId() int64 {
	return r._userId
}
