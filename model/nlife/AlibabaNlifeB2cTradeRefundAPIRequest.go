package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNlifeB2cTradeRefundAPIRequest 零售+请求退款 API请求
// alibaba.nlife.b2c.trade.refund
//
// 零售+平台请求退款接口，在零售+平台不会有资金流变动，只是订单状态的更新
type AlibabaNlifeB2cTradeRefundAPIRequest struct {
	model.Params
	// 退款资金列表
	_refundBillList []FundBill
	// 所退货物的ID列表,逗号分隔商品组，冒号分隔商品和退货数量，支持三种方式退。 条码模式：barcode1:3,barcode2:2 表示barcode1退3件，barcode2退2件。 item_sku模式:itemId1_skuId1:3,itemId2_skuId2:2 表示itemId1_skuId1这个商品退3件，itemId2_skuId2这个商品退2件。 唯一码模式：uniqeueCodeA:1,uniqeueCodeA:1,因唯一码指定到唯一一件商品，退货数量都是1。
	_refundGoodsList []string
	// 零售+平台订单号，和out_trade_no不能同时为空
	_tradeNo string
	// 外部请求号
	_outRequestNo string
	// 外部订单号，和trade_no不能同时为空
	_outTradeNo string
	// 零售+门店ID
	_storeId string
	// 退积分,ISV自行算好
	_refundPoints int64
}

// NewAlibabaNlifeB2cTradeRefundRequest 初始化AlibabaNlifeB2cTradeRefundAPIRequest对象
func NewAlibabaNlifeB2cTradeRefundRequest() *AlibabaNlifeB2cTradeRefundAPIRequest {
	return &AlibabaNlifeB2cTradeRefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaNlifeB2cTradeRefundAPIRequest) GetApiMethodName() string {
	return "alibaba.nlife.b2c.trade.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaNlifeB2cTradeRefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaNlifeB2cTradeRefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundBillList is RefundBillList Setter
// 退款资金列表
func (r *AlibabaNlifeB2cTradeRefundAPIRequest) SetRefundBillList(_refundBillList []FundBill) error {
	r._refundBillList = _refundBillList
	r.Set("refund_bill_list", _refundBillList)
	return nil
}

// GetRefundBillList RefundBillList Getter
func (r AlibabaNlifeB2cTradeRefundAPIRequest) GetRefundBillList() []FundBill {
	return r._refundBillList
}

// SetRefundGoodsList is RefundGoodsList Setter
// 所退货物的ID列表,逗号分隔商品组，冒号分隔商品和退货数量，支持三种方式退。 条码模式：barcode1:3,barcode2:2 表示barcode1退3件，barcode2退2件。 item_sku模式:itemId1_skuId1:3,itemId2_skuId2:2 表示itemId1_skuId1这个商品退3件，itemId2_skuId2这个商品退2件。 唯一码模式：uniqeueCodeA:1,uniqeueCodeA:1,因唯一码指定到唯一一件商品，退货数量都是1。
func (r *AlibabaNlifeB2cTradeRefundAPIRequest) SetRefundGoodsList(_refundGoodsList []string) error {
	r._refundGoodsList = _refundGoodsList
	r.Set("refund_goods_list", _refundGoodsList)
	return nil
}

// GetRefundGoodsList RefundGoodsList Getter
func (r AlibabaNlifeB2cTradeRefundAPIRequest) GetRefundGoodsList() []string {
	return r._refundGoodsList
}

// SetTradeNo is TradeNo Setter
// 零售+平台订单号，和out_trade_no不能同时为空
func (r *AlibabaNlifeB2cTradeRefundAPIRequest) SetTradeNo(_tradeNo string) error {
	r._tradeNo = _tradeNo
	r.Set("trade_no", _tradeNo)
	return nil
}

// GetTradeNo TradeNo Getter
func (r AlibabaNlifeB2cTradeRefundAPIRequest) GetTradeNo() string {
	return r._tradeNo
}

// SetOutRequestNo is OutRequestNo Setter
// 外部请求号
func (r *AlibabaNlifeB2cTradeRefundAPIRequest) SetOutRequestNo(_outRequestNo string) error {
	r._outRequestNo = _outRequestNo
	r.Set("out_request_no", _outRequestNo)
	return nil
}

// GetOutRequestNo OutRequestNo Getter
func (r AlibabaNlifeB2cTradeRefundAPIRequest) GetOutRequestNo() string {
	return r._outRequestNo
}

// SetOutTradeNo is OutTradeNo Setter
// 外部订单号，和trade_no不能同时为空
func (r *AlibabaNlifeB2cTradeRefundAPIRequest) SetOutTradeNo(_outTradeNo string) error {
	r._outTradeNo = _outTradeNo
	r.Set("out_trade_no", _outTradeNo)
	return nil
}

// GetOutTradeNo OutTradeNo Getter
func (r AlibabaNlifeB2cTradeRefundAPIRequest) GetOutTradeNo() string {
	return r._outTradeNo
}

// SetStoreId is StoreId Setter
// 零售+门店ID
func (r *AlibabaNlifeB2cTradeRefundAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabaNlifeB2cTradeRefundAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetRefundPoints is RefundPoints Setter
// 退积分,ISV自行算好
func (r *AlibabaNlifeB2cTradeRefundAPIRequest) SetRefundPoints(_refundPoints int64) error {
	r._refundPoints = _refundPoints
	r.Set("refund_points", _refundPoints)
	return nil
}

// GetRefundPoints RefundPoints Getter
func (r AlibabaNlifeB2cTradeRefundAPIRequest) GetRefundPoints() int64 {
	return r._refundPoints
}
