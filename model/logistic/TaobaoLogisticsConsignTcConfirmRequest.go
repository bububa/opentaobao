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
type TaobaoLogisticsConsignTcConfirmRequest struct {
    model.Params
    // 货主id
    _providerId   int64
    // ERP的名称
    _appName   string
    // 扩展字段 K:V;
    _extendFields   string
    // 发货的包裹
    _goodsList   []ConfirmConsignGoodsDTO
    // 已发货的外部订单号
    _outTradeNo   string
    // 卖家id
    _sellerId   int64
    // 待发货的淘宝主交易号
    _tradeId   int64
}

// 初始化TaobaoLogisticsConsignTcConfirmRequest对象
func NewTaobaoLogisticsConsignTcConfirmRequest() *TaobaoLogisticsConsignTcConfirmRequest{
    return &TaobaoLogisticsConsignTcConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsConsignTcConfirmRequest) GetApiMethodName() string {
    return "taobao.logistics.consign.tc.confirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsConsignTcConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProviderId Setter
// 货主id
func (r *TaobaoLogisticsConsignTcConfirmRequest) SetProviderId(_providerId int64) error {
    r._providerId = _providerId
    r.Set("provider_id", _providerId)
    return nil
}

// ProviderId Getter
func (r TaobaoLogisticsConsignTcConfirmRequest) GetProviderId() int64 {
    return r._providerId
}
// AppName Setter
// ERP的名称
func (r *TaobaoLogisticsConsignTcConfirmRequest) SetAppName(_appName string) error {
    r._appName = _appName
    r.Set("app_name", _appName)
    return nil
}

// AppName Getter
func (r TaobaoLogisticsConsignTcConfirmRequest) GetAppName() string {
    return r._appName
}
// ExtendFields Setter
// 扩展字段 K:V;
func (r *TaobaoLogisticsConsignTcConfirmRequest) SetExtendFields(_extendFields string) error {
    r._extendFields = _extendFields
    r.Set("extend_fields", _extendFields)
    return nil
}

// ExtendFields Getter
func (r TaobaoLogisticsConsignTcConfirmRequest) GetExtendFields() string {
    return r._extendFields
}
// GoodsList Setter
// 发货的包裹
func (r *TaobaoLogisticsConsignTcConfirmRequest) SetGoodsList(_goodsList []ConfirmConsignGoodsDTO) error {
    r._goodsList = _goodsList
    r.Set("goods_list", _goodsList)
    return nil
}

// GoodsList Getter
func (r TaobaoLogisticsConsignTcConfirmRequest) GetGoodsList() []ConfirmConsignGoodsDTO {
    return r._goodsList
}
// OutTradeNo Setter
// 已发货的外部订单号
func (r *TaobaoLogisticsConsignTcConfirmRequest) SetOutTradeNo(_outTradeNo string) error {
    r._outTradeNo = _outTradeNo
    r.Set("out_trade_no", _outTradeNo)
    return nil
}

// OutTradeNo Getter
func (r TaobaoLogisticsConsignTcConfirmRequest) GetOutTradeNo() string {
    return r._outTradeNo
}
// SellerId Setter
// 卖家id
func (r *TaobaoLogisticsConsignTcConfirmRequest) SetSellerId(_sellerId int64) error {
    r._sellerId = _sellerId
    r.Set("seller_id", _sellerId)
    return nil
}

// SellerId Getter
func (r TaobaoLogisticsConsignTcConfirmRequest) GetSellerId() int64 {
    return r._sellerId
}
// TradeId Setter
// 待发货的淘宝主交易号
func (r *TaobaoLogisticsConsignTcConfirmRequest) SetTradeId(_tradeId int64) error {
    r._tradeId = _tradeId
    r.Set("trade_id", _tradeId)
    return nil
}

// TradeId Getter
func (r TaobaoLogisticsConsignTcConfirmRequest) GetTradeId() int64 {
    return r._tradeId
}
