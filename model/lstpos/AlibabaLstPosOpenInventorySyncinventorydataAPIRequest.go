package lstpos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstPosOpenInventorySyncinventorydataAPIRequest 商品库存修改同步接口(最多20条库存信息) API请求
// alibaba.lst.pos.open.inventory.syncinventorydata
//
// 商品库存修改同步接口(最多20条库存信息)
type AlibabaLstPosOpenInventorySyncinventorydataAPIRequest struct {
	model.Params
	// 库存对象列表
	_inventoryDTOList []InventoryDto
	// 门店对应的主账号id
	_userId int64
}

// NewAlibabaLstPosOpenInventorySyncinventorydataRequest 初始化AlibabaLstPosOpenInventorySyncinventorydataAPIRequest对象
func NewAlibabaLstPosOpenInventorySyncinventorydataRequest() *AlibabaLstPosOpenInventorySyncinventorydataAPIRequest {
	return &AlibabaLstPosOpenInventorySyncinventorydataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstPosOpenInventorySyncinventorydataAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.pos.open.inventory.syncinventorydata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstPosOpenInventorySyncinventorydataAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is InventoryDTOList Setter
// 库存对象列表
func (r *AlibabaLstPosOpenInventorySyncinventorydataAPIRequest) SetInventoryDTOList(_inventoryDTOList []InventoryDto) error {
	r._inventoryDTOList = _inventoryDTOList
	r.Set("inventory_d_t_o_list", _inventoryDTOList)
	return nil
}

// Get InventoryDTOList Getter
func (r AlibabaLstPosOpenInventorySyncinventorydataAPIRequest) GetInventoryDTOList() []InventoryDto {
	return r._inventoryDTOList
}

// Set is UserId Setter
// 门店对应的主账号id
func (r *AlibabaLstPosOpenInventorySyncinventorydataAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// Get UserId Getter
func (r AlibabaLstPosOpenInventorySyncinventorydataAPIRequest) GetUserId() int64 {
	return r._userId
}
