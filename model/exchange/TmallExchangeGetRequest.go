package exchange

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单笔换货详情 APIRequest
tmall.exchange.get

获取单笔换货详情
*/
type TmallExchangeGetRequest struct {
    model.Params

    // 换货单号ID
    disputeId   int64 

    // 返回字段。目前支持dispute_id, bizorder_id, num, buyer_nick, status, created, modified, reason, title, buyer_logistic_no, seller_logistic_no, bought_sku, exchange_sku, buyer_address, address, buyer_phone, buyer_logistic_name, seller_logistic_name, alipay_no, buyer_name, seller_nick
    fields   []string 

}

func NewTmallExchangeGetRequest() *TmallExchangeGetRequest{
    return &TmallExchangeGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallExchangeGetRequest) GetApiMethodName() string {
    return "tmall.exchange.get"
}

func (r TmallExchangeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallExchangeGetRequest) SetDisputeId(disputeId int64) error {
    r.disputeId = disputeId
    r.Set("dispute_id", disputeId)
    return nil
}

func (r TmallExchangeGetRequest) GetDisputeId() int64 {
    return r.disputeId
}

func (r *TmallExchangeGetRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TmallExchangeGetRequest) GetFields() []string {
    return r.fields
}

