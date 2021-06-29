package nlife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售+平台支付订单 API请求
alibaba.nlife.b2c.trade.pay

零售+平台支付接口，外部商户调用此接口告知零售+支付结果，保持订单状态同步
*/
type AlibabaNlifeB2cTradePayRequest struct {
    model.Params
    // 零售+平台订单号，和out_trade_no不能同时为空
    _tradeNo   string
    // 提货方式：    LOGISTICS("物流发货"),     SELF_DELIVERY("门店自提");
    _pickingUp   string
    // 收货人
    _consignee   string
    // 收货人电话
    _consigneePhoneNum   string
    // 收货人地址
    _consigneeAddress   string
    // ISV处支付时间
    _gmtPayment   string
    // 支付资金各渠道列表
    _fundBillList   []FundBill
    // 外部订单号，和trade_no不能同时为空
    _outTradeNo   string
    // 实付金额，单位人民币分；该字段实际为必选，为兼容已经接入的isv设置成可选
    _actualPayFee   int64
    // 只传out_trade_no时候，零售+门店号一定要传
    _storeId   string
}

// 初始化AlibabaNlifeB2cTradePayRequest对象
func NewAlibabaNlifeB2cTradePayRequest() *AlibabaNlifeB2cTradePayRequest{
    return &AlibabaNlifeB2cTradePayRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaNlifeB2cTradePayRequest) GetApiMethodName() string {
    return "alibaba.nlife.b2c.trade.pay"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaNlifeB2cTradePayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TradeNo Setter
// 零售+平台订单号，和out_trade_no不能同时为空
func (r *AlibabaNlifeB2cTradePayRequest) SetTradeNo(_tradeNo string) error {
    r._tradeNo = _tradeNo
    r.Set("trade_no", _tradeNo)
    return nil
}

// TradeNo Getter
func (r AlibabaNlifeB2cTradePayRequest) GetTradeNo() string {
    return r._tradeNo
}
// PickingUp Setter
// 提货方式：    LOGISTICS("物流发货"),     SELF_DELIVERY("门店自提");
func (r *AlibabaNlifeB2cTradePayRequest) SetPickingUp(_pickingUp string) error {
    r._pickingUp = _pickingUp
    r.Set("picking_up", _pickingUp)
    return nil
}

// PickingUp Getter
func (r AlibabaNlifeB2cTradePayRequest) GetPickingUp() string {
    return r._pickingUp
}
// Consignee Setter
// 收货人
func (r *AlibabaNlifeB2cTradePayRequest) SetConsignee(_consignee string) error {
    r._consignee = _consignee
    r.Set("consignee", _consignee)
    return nil
}

// Consignee Getter
func (r AlibabaNlifeB2cTradePayRequest) GetConsignee() string {
    return r._consignee
}
// ConsigneePhoneNum Setter
// 收货人电话
func (r *AlibabaNlifeB2cTradePayRequest) SetConsigneePhoneNum(_consigneePhoneNum string) error {
    r._consigneePhoneNum = _consigneePhoneNum
    r.Set("consignee_phone_num", _consigneePhoneNum)
    return nil
}

// ConsigneePhoneNum Getter
func (r AlibabaNlifeB2cTradePayRequest) GetConsigneePhoneNum() string {
    return r._consigneePhoneNum
}
// ConsigneeAddress Setter
// 收货人地址
func (r *AlibabaNlifeB2cTradePayRequest) SetConsigneeAddress(_consigneeAddress string) error {
    r._consigneeAddress = _consigneeAddress
    r.Set("consignee_address", _consigneeAddress)
    return nil
}

// ConsigneeAddress Getter
func (r AlibabaNlifeB2cTradePayRequest) GetConsigneeAddress() string {
    return r._consigneeAddress
}
// GmtPayment Setter
// ISV处支付时间
func (r *AlibabaNlifeB2cTradePayRequest) SetGmtPayment(_gmtPayment string) error {
    r._gmtPayment = _gmtPayment
    r.Set("gmt_payment", _gmtPayment)
    return nil
}

// GmtPayment Getter
func (r AlibabaNlifeB2cTradePayRequest) GetGmtPayment() string {
    return r._gmtPayment
}
// FundBillList Setter
// 支付资金各渠道列表
func (r *AlibabaNlifeB2cTradePayRequest) SetFundBillList(_fundBillList []FundBill) error {
    r._fundBillList = _fundBillList
    r.Set("fund_bill_list", _fundBillList)
    return nil
}

// FundBillList Getter
func (r AlibabaNlifeB2cTradePayRequest) GetFundBillList() []FundBill {
    return r._fundBillList
}
// OutTradeNo Setter
// 外部订单号，和trade_no不能同时为空
func (r *AlibabaNlifeB2cTradePayRequest) SetOutTradeNo(_outTradeNo string) error {
    r._outTradeNo = _outTradeNo
    r.Set("out_trade_no", _outTradeNo)
    return nil
}

// OutTradeNo Getter
func (r AlibabaNlifeB2cTradePayRequest) GetOutTradeNo() string {
    return r._outTradeNo
}
// ActualPayFee Setter
// 实付金额，单位人民币分；该字段实际为必选，为兼容已经接入的isv设置成可选
func (r *AlibabaNlifeB2cTradePayRequest) SetActualPayFee(_actualPayFee int64) error {
    r._actualPayFee = _actualPayFee
    r.Set("actual_pay_fee", _actualPayFee)
    return nil
}

// ActualPayFee Getter
func (r AlibabaNlifeB2cTradePayRequest) GetActualPayFee() int64 {
    return r._actualPayFee
}
// StoreId Setter
// 只传out_trade_no时候，零售+门店号一定要传
func (r *AlibabaNlifeB2cTradePayRequest) SetStoreId(_storeId string) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r AlibabaNlifeB2cTradePayRequest) GetStoreId() string {
    return r._storeId
}
