package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoInventoryInitialItemAPIRequest
商品库存初始化 API请求
taobao.inventory.initial.item

建议使用新接口：taobao.inventory.merchant.adjust ，该接口会逐步停用。
商家仓商品初始化在各个仓中库存 */
type TaobaoInventoryInitialItemAPIRequest struct {
	model.Params
	// 后端商品id
	_scItemId int64
	// 商品初始库存信息： [{"storeCode":"必选,商家仓库编号","inventoryType":"可选，库存类型 1：正常,2：损坏,3：冻结,10：质押,11-20:用户自定义,默认为1","quantity":"必选,数量"}]
	_storeInventorys string
}

// New
