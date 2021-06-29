package nlife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售+平台支付订单 APIRequest
alibaba.nlife.b2c.trade.pay

零售+平台支付接口，外部商户调用此接口告知零售+支付结果，保持订单状态同步
*/
type AlibabaNlifeB2cTradePayRequest struct {
    model.Params

    // 零售+平台订单号，和out_trade_no不能同时为空
    tradeNo   string 

    // 提货方式：    LOGISTICS("物流发货"),     SELF_DELIVERY("门店自提");
    pickingUp   string 

    // 收货人
    consignee   string 

    // 收货人电话
    consigneePhoneNum   string 

    // 收货人地址
    consigneeAddress   string 

    // ISV处支付时间
    gmtPayment   string 

    // 支付资金各渠道列表
    fundBillList   []FundBill 

    // 外部订单号，和trade_no不能同时为空
    outTradeNo   string 

    // 实付金额，单位人民币分；该字段实际为必选，为兼容已经接入的isv设置成可选
    actualPayFee   int64 

    // 只传out_trade_no时候，零售+门店号一定要传
    storeId   string 

}

func NewAlibabaNlifeB2cTradePayRequest() *AlibabaNlifeB2cTradePayRequest{
    return &AlibabaNlifeB2cTradePayRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaNlifeB2cTradePayRequest) GetApiMethodName() string {
    return "alibaba.nlife.b2c.trade.pay"
}

func (r AlibabaNlifeB2cTradePayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaNlifeB2cTradePayRequest) SetTradeNo(tradeNo string) error {
    r.tradeNo = tradeNo
    r.Set("trade_no", tradeNo)
    return nil
}

func (r AlibabaNlifeB2cTradePayRequest) GetTradeNo() string {
    return r.tradeNo
}

func (r *AlibabaNlifeB2cTradePayRequest) SetPickingUp(pickingUp string) error {
    r.pickingUp = pickingUp
    r.Set("picking_up", pickingUp)
    return nil
}

func (r AlibabaNlifeB2cTradePayRequest) GetPickingUp() string {
    return r.pickingUp
}

func (r *AlibabaNlifeB2cTradePayRequest) SetConsignee(consignee string) error {
    r.consignee = consignee
    r.Set("consignee", consignee)
    return nil
}

func (r AlibabaNlifeB2cTradePayRequest) GetConsignee() string {
    return r.consignee
}

func (r *AlibabaNlifeB2cTradePayRequest) SetConsigneePhoneNum(consigneePhoneNum string) error {
    r.consigneePhoneNum = consigneePhoneNum
    r.Set("consignee_phone_num", consigneePhoneNum)
    return nil
}

func (r AlibabaNlifeB2cTradePayRequest) GetConsigneePhoneNum() string {
    return r.consigneePhoneNum
}

func (r *AlibabaNlifeB2cTradePayRequest) SetConsigneeAddress(consigneeAddress string) error {
    r.consigneeAddress = consigneeAddress
    r.Set("consignee_address", consigneeAddress)
    return nil
}

func (r AlibabaNlifeB2cTradePayRequest) GetConsigneeAddress() string {
    return r.consigneeAddress
}

func (r *AlibabaNlifeB2cTradePayRequest) SetGmtPayment(gmtPayment string) error {
    r.gmtPayment = gmtPayment
    r.Set("gmt_payment", gmtPayment)
    return nil
}

func (r AlibabaNlifeB2cTradePayRequest) GetGmtPayment() string {
    return r.gmtPayment
}

func (r *AlibabaNlifeB2cTradePayRequest) SetFundBillList(fundBillList []FundBill) error {
    r.fundBillList = fundBillList
    r.Set("fund_bill_list", fundBillList)
    return nil
}

func (r AlibabaNlifeB2cTradePayRequest) GetFundBillList() []FundBill {
    return r.fundBillList
}

func (r *AlibabaNlifeB2cTradePayRequest) SetOutTradeNo(outTradeNo string) error {
    r.outTradeNo = outTradeNo
    r.Set("out_trade_no", outTradeNo)
    return nil
}

func (r AlibabaNlifeB2cTradePayRequest) GetOutTradeNo() string {
    return r.outTradeNo
}

func (r *AlibabaNlifeB2cTradePayRequest) SetActualPayFee(actualPayFee int64) error {
    r.actualPayFee = actualPayFee
    r.Set("actual_pay_fee", actualPayFee)
    return nil
}

func (r AlibabaNlifeB2cTradePayRequest) GetActualPayFee() int64 {
    return r.actualPayFee
}

func (r *AlibabaNlifeB2cTradePayRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r AlibabaNlifeB2cTradePayRequest) GetStoreId() string {
    return r.storeId
}

