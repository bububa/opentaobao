package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbInventoryDetailGetAPIRequest
查询库存详情 API请求
taobao.wlb.inventory.detail.get

查询库存详情，通过商品ID获取发送请求的卖家的库存详情 */
type TaobaoWlbInventoryDetailGetAPIRequest struct {
	model.Params
	// 库存类型列表，值包括：<br/>VENDIBLE--可销售库存<br/>FREEZE--冻结库存<br/>ONWAY--在途库存<br/>DEFECT--残次品库存<br/>ENGINE_DAMAGE--机损<br/>BOX_DAMAGE--箱损<br/>EXPIRATION--过保
	_inventoryTypeList []string
	// 仓库编码
	_storeCode string
	// 商品ID
	_itemId int64
}

// New
