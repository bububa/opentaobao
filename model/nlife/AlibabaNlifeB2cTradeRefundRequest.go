package nlife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售+请求退款 API请求
alibaba.nlife.b2c.trade.refund

零售+平台请求退款接口，在零售+平台不会有资金流变动，只是订单状态的更新
*/
type AlibabaNlifeB2cTradeRefundRequest struct {
    model.Params
    // 零售+平台订单号，和out_trade_no不能同时为空
    _tradeNo   string
    // 外部请求号
    _outRequestNo   string
    // 退款资金列表
    _refundBillList   []FundBill
    // 所退货物的ID列表,逗号分隔商品组，冒号分隔商品和退货数量，支持三种方式退。 条码模式：barcode1:3,barcode2:2 表示barcode1退3件，barcode2退2件。 item_sku模式:itemId1_skuId1:3,itemId2_skuId2:2 表示itemId1_skuId1这个商品退3件，itemId2_skuId2这个商品退2件。 唯一码模式：uniqeueCodeA:1,uniqeueCodeA:1,因唯一码指定到唯一一件商品，退货数量都是1。
    _refundGoodsList   []string
    // 外部订单号，和trade_no不能同时为空
    _outTradeNo   string
    // 零售+门店ID
    _storeId   string
    // 退积分,ISV自行算好
    _refundPoints   int64
}

// 初始化AlibabaNlifeB2cTradeRefundRequest对象
func NewAlibabaNlifeB2cTradeRefundRequest() *AlibabaNlifeB2cTradeRefundRequest{
    return &AlibabaNlifeB2cTradeRefundRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaNlifeB2cTradeRefundRequest) GetApiMethodName() string {
    return "alibaba.nlife.b2c.trade.refund"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaNlifeB2cTradeRefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TradeNo Setter
// 零售+平台订单号，和out_trade_no不能同时为空
func (r *AlibabaNlifeB2cTradeRefundRequest) SetTradeNo(_tradeNo string) error {
    r._tradeNo = _tradeNo
    r.Set("trade_no", _tradeNo)
    return nil
}

// TradeNo Getter
func (r AlibabaNlifeB2cTradeRefundRequest) GetTradeNo() string {
    return r._tradeNo
}
// OutRequestNo Setter
// 外部请求号
func (r *AlibabaNlifeB2cTradeRefundRequest) SetOutRequestNo(_outRequestNo string) error {
    r._outRequestNo = _outRequestNo
    r.Set("out_request_no", _outRequestNo)
    return nil
}

// OutRequestNo Getter
func (r AlibabaNlifeB2cTradeRefundRequest) GetOutRequestNo() string {
    return r._outRequestNo
}
// RefundBillList Setter
// 退款资金列表
func (r *AlibabaNlifeB2cTradeRefundRequest) SetRefundBillList(_refundBillList []FundBill) error {
    r._refundBillList = _refundBillList
    r.Set("refund_bill_list", _refundBillList)
    return nil
}

// RefundBillList Getter
func (r AlibabaNlifeB2cTradeRefundRequest) GetRefundBillList() []FundBill {
    return r._refundBillList
}
// RefundGoodsList Setter
// 所退货物的ID列表,逗号分隔商品组，冒号分隔商品和退货数量，支持三种方式退。 条码模式：barcode1:3,barcode2:2 表示barcode1退3件，barcode2退2件。 item_sku模式:itemId1_skuId1:3,itemId2_skuId2:2 表示itemId1_skuId1这个商品退3件，itemId2_skuId2这个商品退2件。 唯一码模式：uniqeueCodeA:1,uniqeueCodeA:1,因唯一码指定到唯一一件商品，退货数量都是1。
func (r *AlibabaNlifeB2cTradeRefundRequest) SetRefundGoodsList(_refundGoodsList []string) error {
    r._refundGoodsList = _refundGoodsList
    r.Set("refund_goods_list", _refundGoodsList)
    return nil
}

// RefundGoodsList Getter
func (r AlibabaNlifeB2cTradeRefundRequest) GetRefundGoodsList() []string {
    return r._refundGoodsList
}
// OutTradeNo Setter
// 外部订单号，和trade_no不能同时为空
func (r *AlibabaNlifeB2cTradeRefundRequest) SetOutTradeNo(_outTradeNo string) error {
    r._outTradeNo = _outTradeNo
    r.Set("out_trade_no", _outTradeNo)
    return nil
}

// OutTradeNo Getter
func (r AlibabaNlifeB2cTradeRefundRequest) GetOutTradeNo() string {
    return r._outTradeNo
}
// StoreId Setter
// 零售+门店ID
func (r *AlibabaNlifeB2cTradeRefundRequest) SetStoreId(_storeId string) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r AlibabaNlifeB2cTradeRefundRequest) GetStoreId() string {
    return r._storeId
}
// RefundPoints Setter
// 退积分,ISV自行算好
func (r *AlibabaNlifeB2cTradeRefundRequest) SetRefundPoints(_refundPoints int64) error {
    r._refundPoints = _refundPoints
    r.Set("refund_points", _refundPoints)
    return nil
}

// RefundPoints Getter
func (r AlibabaNlifeB2cTradeRefundRequest) GetRefundPoints() int64 {
    return r._refundPoints
}
