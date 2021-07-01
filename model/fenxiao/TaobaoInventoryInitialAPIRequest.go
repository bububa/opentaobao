package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoInventoryInitialAPIRequest
库存初始化 API请求
taobao.inventory.initial

建议使用新接口：taobao.inventory.merchant.adjust ，该接口会逐步停用。
商家仓库存初始化接口，直接按照商家指定的商品库存数进行填充，没有单据核对，不参与库存对账。 */
type TaobaoInventoryInitialAPIRequest struct {
	model.Params
	// 商家仓库编码
	_storeCode string
	// 商品初始库存信息： [{"scItemId":"商品后端ID，如果有传scItemCode,参数可以为0","scItemCode":"商品商家编码","inventoryType":"库存类型  1：正常,2：损坏,3：冻结,10：质押,11-20:用户自定义","quantity":"数量"}]
	_items string
}

// New
