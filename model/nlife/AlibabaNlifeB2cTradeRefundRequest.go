package nlife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售+请求退款 APIRequest
alibaba.nlife.b2c.trade.refund

零售+平台请求退款接口，在零售+平台不会有资金流变动，只是订单状态的更新
*/
type AlibabaNlifeB2cTradeRefundRequest struct {
    model.Params

    // 零售+平台订单号，和out_trade_no不能同时为空
    tradeNo   string 

    // 外部请求号
    outRequestNo   string 

    // 退款资金列表
    refundBillList   []FundBill 

    // 所退货物的ID列表,逗号分隔商品组，冒号分隔商品和退货数量，支持三种方式退。 条码模式：barcode1:3,barcode2:2 表示barcode1退3件，barcode2退2件。 item_sku模式:itemId1_skuId1:3,itemId2_skuId2:2 表示itemId1_skuId1这个商品退3件，itemId2_skuId2这个商品退2件。 唯一码模式：uniqeueCodeA:1,uniqeueCodeA:1,因唯一码指定到唯一一件商品，退货数量都是1。
    refundGoodsList   []string 

    // 外部订单号，和trade_no不能同时为空
    outTradeNo   string 

    // 零售+门店ID
    storeId   string 

    // 退积分,ISV自行算好
    refundPoints   int64 

}

func NewAlibabaNlifeB2cTradeRefundRequest() *AlibabaNlifeB2cTradeRefundRequest{
    return &AlibabaNlifeB2cTradeRefundRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaNlifeB2cTradeRefundRequest) GetApiMethodName() string {
    return "alibaba.nlife.b2c.trade.refund"
}

func (r AlibabaNlifeB2cTradeRefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaNlifeB2cTradeRefundRequest) SetTradeNo(tradeNo string) error {
    r.tradeNo = tradeNo
    r.Set("trade_no", tradeNo)
    return nil
}

func (r AlibabaNlifeB2cTradeRefundRequest) GetTradeNo() string {
    return r.tradeNo
}

func (r *AlibabaNlifeB2cTradeRefundRequest) SetOutRequestNo(outRequestNo string) error {
    r.outRequestNo = outRequestNo
    r.Set("out_request_no", outRequestNo)
    return nil
}

func (r AlibabaNlifeB2cTradeRefundRequest) GetOutRequestNo() string {
    return r.outRequestNo
}

func (r *AlibabaNlifeB2cTradeRefundRequest) SetRefundBillList(refundBillList []FundBill) error {
    r.refundBillList = refundBillList
    r.Set("refund_bill_list", refundBillList)
    return nil
}

func (r AlibabaNlifeB2cTradeRefundRequest) GetRefundBillList() []FundBill {
    return r.refundBillList
}

func (r *AlibabaNlifeB2cTradeRefundRequest) SetRefundGoodsList(refundGoodsList []string) error {
    r.refundGoodsList = refundGoodsList
    r.Set("refund_goods_list", refundGoodsList)
    return nil
}

func (r AlibabaNlifeB2cTradeRefundRequest) GetRefundGoodsList() []string {
    return r.refundGoodsList
}

func (r *AlibabaNlifeB2cTradeRefundRequest) SetOutTradeNo(outTradeNo string) error {
    r.outTradeNo = outTradeNo
    r.Set("out_trade_no", outTradeNo)
    return nil
}

func (r AlibabaNlifeB2cTradeRefundRequest) GetOutTradeNo() string {
    return r.outTradeNo
}

func (r *AlibabaNlifeB2cTradeRefundRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r AlibabaNlifeB2cTradeRefundRequest) GetStoreId() string {
    return r.storeId
}

func (r *AlibabaNlifeB2cTradeRefundRequest) SetRefundPoints(refundPoints int64) error {
    r.refundPoints = refundPoints
    r.Set("refund_points", refundPoints)
    return nil
}

func (r AlibabaNlifeB2cTradeRefundRequest) GetRefundPoints() int64 {
    return r.refundPoints
}

