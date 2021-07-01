package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoInventoryAdjustTradeAPIRequest
交易库存调整单 API请求
taobao.inventory.adjust.trade

商家交易调整库存，淘宝交易、B2B经销等 */
type TaobaoInventoryAdjustTradeAPIRequest struct {
	model.Params
	// 订单类型：B2C、B2B
	_tbOrderType string
	// 业务操作时间
	_operateTime string
	// 商家外部定单号
	_bizUniqueCode string
	// 商品初始库存信息： [{ "TBOrderCode”:”淘宝交易号”,"TBSubOrderCode ":"淘宝子交易单号,赠品可以不填","”isGift”:”TRUE或者FALSE,是否赠品”,storeCode":"商家仓库编码"," scItemId ":"商品后端ID","scItemCode":"商品商家编码"," originScItemId ":"原商品ID","inventoryType":"","quantity":"111","isComplete":"TRUE或者FALSE，是否全部确认出库"}]
	_items string
}

// New
