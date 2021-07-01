package lstpos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstPosOpenInventorySyncinventorydataAPIRequest
商品库存修改同步接口(最多20条库存信息) API请求
alibaba.lst.pos.open.inventory.syncinventorydata

商品库存修改同步接口(最多20条库存信息) */
type AlibabaLstPosOpenInventorySyncinventorydataAPIRequest struct {
	model.Params
	// 库存对象列表
	_inventoryDTOList []InventoryDto
	// 门店对应的主账号id
	_userId int64
}

// New
