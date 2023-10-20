package lstpos

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstPosOpenInventorySyncinventorydataAPIRequest) Reset() {
	r._inventoryDTOList = r._inventoryDTOList[:0]
	r._userId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstPosOpenInventorySyncinventorydataAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.pos.open.inventory.syncinventorydata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstPosOpenInventorySyncinventorydataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstPosOpenInventorySyncinventorydataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInventoryDTOList is InventoryDTOList Setter
// 库存对象列表
func (r *AlibabaLstPosOpenInventorySyncinventorydataAPIRequest) SetInventoryDTOList(_inventoryDTOList []InventoryDto) error {
	r._inventoryDTOList = _inventoryDTOList
	r.Set("inventory_d_t_o_list", _inventoryDTOList)
	return nil
}

// GetInventoryDTOList InventoryDTOList Getter
func (r AlibabaLstPosOpenInventorySyncinventorydataAPIRequest) GetInventoryDTOList() []InventoryDto {
	return r._inventoryDTOList
}

// SetUserId is UserId Setter
// 门店对应的主账号id
func (r *AlibabaLstPosOpenInventorySyncinventorydataAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaLstPosOpenInventorySyncinventorydataAPIRequest) GetUserId() int64 {
	return r._userId
}

var poolAlibabaLstPosOpenInventorySyncinventorydataAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstPosOpenInventorySyncinventorydataRequest()
	},
}

// GetAlibabaLstPosOpenInventorySyncinventorydataRequest 从 sync.Pool 获取 AlibabaLstPosOpenInventorySyncinventorydataAPIRequest
func GetAlibabaLstPosOpenInventorySyncinventorydataAPIRequest() *AlibabaLstPosOpenInventorySyncinventorydataAPIRequest {
	return poolAlibabaLstPosOpenInventorySyncinventorydataAPIRequest.Get().(*AlibabaLstPosOpenInventorySyncinventorydataAPIRequest)
}

// ReleaseAlibabaLstPosOpenInventorySyncinventorydataAPIRequest 将 AlibabaLstPosOpenInventorySyncinventorydataAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstPosOpenInventorySyncinventorydataAPIRequest(v *AlibabaLstPosOpenInventorySyncinventorydataAPIRequest) {
	v.Reset()
	poolAlibabaLstPosOpenInventorySyncinventorydataAPIRequest.Put(v)
}
