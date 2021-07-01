package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoInventoryAdjustExternalAPIRequest
非交易库存调整单 API请求
taobao.inventory.adjust.external

建议使用新接口：taobao.inventory.merchant.adjust ，该接口会逐步停用。
商家非交易调整库存，调拨出库、盘点等时调用 */
type TaobaoInventoryAdjustExternalAPIRequest struct {
	model.Params
	// test
	_reduceType string
	// 外部订单类型, BALANCE:盘点、NON_TAOBAO_TRADE:非淘宝交易、ALLOCATE:调拨、OTHERS:其他
	_bizType string
	// test
	_operateType string
	// 商家外部定单号
	_bizUniqueCode string
	// 商品初始库存信息： [{"scItemId":"商品后端ID，如果有传scItemCode,参数可以为0","scItemCode":"商品商家编码","inventoryType":"库存类型  1：正常,”direction”: 1: 盘盈 -1: 盘亏,参数可选,"quantity":"数量(正数)"}]
	_items string
	// 库存占用返回的操作码.operate_type 为OUTBOUND时，如果是确认事先进行过的库存占用，需要传入当时返回的操作码，并且明细必须与申请时保持一致
	_occupyOperateCode string
	// 业务操作时间
	_operateTime string
	// 商家仓库编码
	_storeCode string
}

// New
