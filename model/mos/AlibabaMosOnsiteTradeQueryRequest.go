package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新商场当面付交易查询 APIRequest
alibaba.mos.onsite.trade.query

本接口提供新商场当面付订单的查询的功能，商户可以通过本接口主动查询订单状态，完成下一步的业务逻辑。
商户系统应在两种场景下调用此接口: 商户POS系统应该在调用[条码支付请求接口]并成功返回后，调用此接口查询订单的处理状态。
*/
type AlibabaMosOnsiteTradeQueryRequest struct {
    model.Params

    // 喵街交易流水号。与商户支付流水号两者至少要填写一个。如果均有，优先级为喵街交易流水号  > 商户支付流水号。
    tradeNo   string 

    // 商户ID类型，取值为miaojie或out
    storeIdType   string 

    // 门店号或喵街商户ID
    storeId   string 

    // 原支付请求的商户支付流水号。与喵街交易流水号两者至少要填写一个。如果均有，优先级为喵街交易流水号 >  商户支付流水号。
    outTradeNo   string 

}

func NewAlibabaMosOnsiteTradeQueryRequest() *AlibabaMosOnsiteTradeQueryRequest{
    return &AlibabaMosOnsiteTradeQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMosOnsiteTradeQueryRequest) GetApiMethodName() string {
    return "alibaba.mos.onsite.trade.query"
}

func (r AlibabaMosOnsiteTradeQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMosOnsiteTradeQueryRequest) SetTradeNo(tradeNo string) error {
    r.tradeNo = tradeNo
    r.Set("trade_no", tradeNo)
    return nil
}

func (r AlibabaMosOnsiteTradeQueryRequest) GetTradeNo() string {
    return r.tradeNo
}

func (r *AlibabaMosOnsiteTradeQueryRequest) SetStoreIdType(storeIdType string) error {
    r.storeIdType = storeIdType
    r.Set("store_id_type", storeIdType)
    return nil
}

func (r AlibabaMosOnsiteTradeQueryRequest) GetStoreIdType() string {
    return r.storeIdType
}

func (r *AlibabaMosOnsiteTradeQueryRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r AlibabaMosOnsiteTradeQueryRequest) GetStoreId() string {
    return r.storeId
}

func (r *AlibabaMosOnsiteTradeQueryRequest) SetOutTradeNo(outTradeNo string) error {
    r.outTradeNo = outTradeNo
    r.Set("out_trade_no", outTradeNo)
    return nil
}

func (r AlibabaMosOnsiteTradeQueryRequest) GetOutTradeNo() string {
    return r.outTradeNo
}

