package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通知交易确认发货接口 APIRequest
taobao.logistics.consign.tc.confirm

下述业务场景可以使用此接口通知相关的交易订单发货：
1、发货过程分为多段操作，在确定发货之前，不需要通知交易，当货确认已发出之后，才通知交易发货。
2、发货过程涉及到多个订单，其中一个订单是跟实操的发货操作同步的，剩下的订单，需要在实操的订单发货之后，一并通知交易发货。
*/
type TaobaoLogisticsConsignTcConfirmRequest struct {
    model.Params

    // 货主id
    providerId   int64 

    // ERP的名称
    appName   string 

    // 扩展字段 K:V;
    extendFields   string 

    // 发货的包裹
    goodsList   []ConfirmConsignGoodsDto 

    // 已发货的外部订单号
    outTradeNo   string 

    // 卖家id
    sellerId   int64 

    // 待发货的淘宝主交易号
    tradeId   int64 

}

func NewTaobaoLogisticsConsignTcConfirmRequest() *TaobaoLogisticsConsignTcConfirmRequest{
    return &TaobaoLogisticsConsignTcConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoLogisticsConsignTcConfirmRequest) GetApiMethodName() string {
    return "taobao.logistics.consign.tc.confirm"
}

func (r TaobaoLogisticsConsignTcConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoLogisticsConsignTcConfirmRequest) SetProviderId(providerId int64) error {
    r.providerId = providerId
    r.Set("provider_id", providerId)
    return nil
}

func (r TaobaoLogisticsConsignTcConfirmRequest) GetProviderId() int64 {
    return r.providerId
}

func (r *TaobaoLogisticsConsignTcConfirmRequest) SetAppName(appName string) error {
    r.appName = appName
    r.Set("app_name", appName)
    return nil
}

func (r TaobaoLogisticsConsignTcConfirmRequest) GetAppName() string {
    return r.appName
}

func (r *TaobaoLogisticsConsignTcConfirmRequest) SetExtendFields(extendFields string) error {
    r.extendFields = extendFields
    r.Set("extend_fields", extendFields)
    return nil
}

func (r TaobaoLogisticsConsignTcConfirmRequest) GetExtendFields() string {
    return r.extendFields
}

func (r *TaobaoLogisticsConsignTcConfirmRequest) SetGoodsList(goodsList []ConfirmConsignGoodsDto) error {
    r.goodsList = goodsList
    r.Set("goods_list", goodsList)
    return nil
}

func (r TaobaoLogisticsConsignTcConfirmRequest) GetGoodsList() []ConfirmConsignGoodsDto {
    return r.goodsList
}

func (r *TaobaoLogisticsConsignTcConfirmRequest) SetOutTradeNo(outTradeNo string) error {
    r.outTradeNo = outTradeNo
    r.Set("out_trade_no", outTradeNo)
    return nil
}

func (r TaobaoLogisticsConsignTcConfirmRequest) GetOutTradeNo() string {
    return r.outTradeNo
}

func (r *TaobaoLogisticsConsignTcConfirmRequest) SetSellerId(sellerId int64) error {
    r.sellerId = sellerId
    r.Set("seller_id", sellerId)
    return nil
}

func (r TaobaoLogisticsConsignTcConfirmRequest) GetSellerId() int64 {
    return r.sellerId
}

func (r *TaobaoLogisticsConsignTcConfirmRequest) SetTradeId(tradeId int64) error {
    r.tradeId = tradeId
    r.Set("trade_id", tradeId)
    return nil
}

func (r TaobaoLogisticsConsignTcConfirmRequest) GetTradeId() int64 {
    return r.tradeId
}

