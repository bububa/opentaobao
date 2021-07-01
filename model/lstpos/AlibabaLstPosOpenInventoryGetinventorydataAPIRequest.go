package lstpos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstPosOpenInventoryGetinventorydataAPIRequest
商品库存只读接口(最多20条库存信息) API请求
alibaba.lst.pos.open.inventory.getinventorydata

商品库存只读接口(最多20条库存信息) */
type AlibabaLstPosOpenInventoryGetinventorydataAPIRequest struct {
	model.Params
	// ISV商品Id列表
	_isvGoodsIdList []string
	// 门店主账号Id
	_userId int64
}

// New
