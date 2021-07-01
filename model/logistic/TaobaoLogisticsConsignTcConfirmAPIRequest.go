package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通知交易确认发货接口 API请求
taobao.logistics.consign.tc.confirm

下述业务场景可以使用此接口通知相关的交易订单发货：
1、发货过程分为多段操作，在确定发货之前，不需要通知交易，当货确认已发出之后，才通知交易发货。
2、发货过程涉及到多个订单，其中一个订单是跟实操的发货操作同步的，剩下的订单，需要在实操的订单发货之后，一并通知交易发货。
*/
type TaobaoLogisticsConsignTcConfirmAPIRequest struct {
    model.Params
    // 货主id
    _providerId   int64
    // ERP的名称
    _appName   string
    // 扩展字段 K:V;
    _extendFields   string
    // 发货的包裹
    _goodsList   []ConfirmConsignGoodsDto
    // 已发货的外部订单号
    _outTradeNo   string
    // 卖家id
    _sellerId   int64
    // 待发货的淘宝主交易号
    _tradeId   int64
}

// 初始化TaobaoLogisticsConsignTcConfirmAPIRequest对象
func NewTaobaoLogisticsConsignTcConfirmRequest() *TaobaoLogisticsConsignTcConfirmAPIRequest{
    return &TaobaoLogisticsConsignTcConfirmAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsConsignTcConfirmAPIRequest) GetApiMethodName() string {
    return "taobao.logistics.consign.tc.confirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsConsignTcConfirmAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProviderId Setter
// 货主id
func (r *TaobaoLogisticsConsignTcConfirmAPIRequest) SetProviderId(_providerId int64) error {
    r._providerId = _providerId
    r.Set("provider_id", _providerId)
    return nil
}

// ProviderId Getter
func (r TaobaoLogisticsConsignTcConfirmAPIRequest) GetProviderId() int64 {
    return r._providerId
}
// AppName Setter
// ERP的名称
func (r *TaobaoLogisticsConsignTcConfirmAPIRequest) SetAppName(_appName string) error {
    r._appName = _appName
    r.Set("app_name", _appName)
    return nil
}

// AppName Getter
func (r TaobaoLogisticsConsignTcConfirmAPIRequest) GetAppName() string {
    return r._appName
}
// ExtendFields Setter
// 扩展字段 K:V;
func (r *TaobaoLogisticsConsignTcConfirmAPIRequest) SetExtendFields(_extendFields string) error {
    r._extendFields = _extendFields
    r.Set("extend_fields", _extendFields)
    return nil
}

// ExtendFields Getter
func (r TaobaoLogisticsConsignTcConfirmAPIRequest) GetExtendFields() string {
    return r._extendFields
}
// GoodsList Setter
// 发货的包裹
func (r *TaobaoLogisticsConsignTcConfirmAPIRequest) SetGoodsList(_goodsList []ConfirmConsignGoodsDto) error {
    r._goodsList = _goodsList
    r.Set("goods_list", _goodsList)
    return nil
}

// GoodsList Getter
func (r TaobaoLogisticsConsignTcConfirmAPIRequest) GetGoodsList() []ConfirmConsignGoodsDto {
    return r._goodsList
}
// OutTradeNo Setter
// 已发货的外部订单号
func (r *TaobaoLogisticsConsignTcConfirmAPIRequest) SetOutTradeNo(_outTradeNo string) error {
    r._outTradeNo = _outTradeNo
    r.Set("out_trade_no", _outTradeNo)
    return nil
}

// OutTradeNo Getter
func (r TaobaoLogisticsConsignTcConfirmAPIRequest) GetOutTradeNo() string {
    return r._outTradeNo
}
// SellerId Setter
// 卖家id
func (r *TaobaoLogisticsConsignTcConfirmAPIRequest) SetSellerId(_sellerId int64) error {
    r._sellerId = _sellerId
    r.Set("seller_id", _sellerId)
    return nil
}

// SellerId Getter
func (r TaobaoLogisticsConsignTcConfirmAPIRequest) GetSellerId() int64 {
    return r._sellerId
}
// TradeId Setter
// 待发货的淘宝主交易号
func (r *TaobaoLogisticsConsignTcConfirmAPIRequest) SetTradeId(_tradeId int64) error {
    r._tradeId = _tradeId
    r.Set("trade_id", _tradeId)
    return nil
}

// TradeId Getter
func (r TaobaoLogisticsConsignTcConfirmAPIRequest) GetTradeId() int64 {
    return r._tradeId
}
