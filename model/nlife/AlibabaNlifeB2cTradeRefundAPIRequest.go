package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNlifeB2cTradeRefundAPIRequest
零售+请求退款 API请求
alibaba.nlife.b2c.trade.refund

零售+平台请求退款接口，在零售+平台不会有资金流变动，只是订单状态的更新 */
type AlibabaNlifeB2cTradeRefundAPIRequest struct {
	model.Params
	// 零售+平台订单号，和out_trade_no不能同时为空
	_tradeNo string
	// 外部请求号
	_outRequestNo string
	// 退款资金列表
	_refundBillList []FundBill
	// 所退货物的ID列表,逗号分隔商品组，冒号分隔商品和退货数量，支持三种方式退。 条码模式：barcode1:3,barcode2:2 表示barcode1退3件，barcode2退2件。 item_sku模式:itemId1_skuId1:3,itemId2_skuId2:2 表示itemId1_skuId1这个商品退3件，itemId2_skuId2这个商品退2件。 唯一码模式：uniqeueCodeA:1,uniqeueCodeA:1,因唯一码指定到唯一一件商品，退货数量都是1。
	_refundGoodsList []string
	// 外部订单号，和trade_no不能同时为空
	_outTradeNo string
	// 零售+门店ID
	_storeId string
	// 退积分,ISV自行算好
	_refundPoints int64
}

// New
