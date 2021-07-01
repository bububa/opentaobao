package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbOutInventoryChangeNotifyAPIRequest
外部库存变化通知（企业物流用户使用） API请求
taobao.wlb.out.inventory.change.notify

拥有自有仓的企业物流用户通过该接口把自有仓的库存通知到物流宝，由物流宝维护该库存，控制前台显示库存的准确性。 */
type TaobaoWlbOutInventoryChangeNotifyAPIRequest struct {
	model.Params
	// WLB_ITEM--物流宝商品 IC_ITEM--淘宝商品 IC_SKU--淘宝sku
	_type string
	// OUT--出库 IN--入库
	_opType string
	// （1）OTHER： 其他 （2）TAOBAO_TRADE： 淘宝交易 （3）OTHER_TRADE：其他交易 （4）ALLCOATE： 调拨 （5）CHECK:盘点 （6）PURCHASE:采购
	_source string
	// 物流宝商品id或前台宝贝id（由type类型决定）
	_itemId int64
	// 库存变化数量
	_changeCount int64
	// 本次库存变化后库存余额
	_resultCount int64
	// 订单号，如果source为TAOBAO_TRADE,order_source_code必须为淘宝交易号
	_orderSourceCode string
	// 库存变化唯一标识，用于去重，一个外部唯一编码唯一标识一次库存变化。
	_outBizCode string
	// 目前非必须，系统会选择默认值
	_storeCode string
}

// New
