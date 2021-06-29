package exchange

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单笔换货详情 API请求
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

// 初始化TmallExchangeGetRequest对象
func NewTmallExchangeGetRequest() *TmallExchangeGetRequest{
    return &TmallExchangeGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallExchangeGetRequest) GetApiMethodName() string {
    return "tmall.exchange.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallExchangeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DisputeId Setter
// 换货单号ID
func (r *TmallExchangeGetRequest) SetDisputeId(disputeId int64) error {
    r.disputeId = disputeId
    r.Set("dispute_id", disputeId)
    return nil
}

// DisputeId Getter
func (r TmallExchangeGetRequest) GetDisputeId() int64 {
    return r.disputeId
}
// Fields Setter
// 返回字段。目前支持dispute_id, bizorder_id, num, buyer_nick, status, created, modified, reason, title, buyer_logistic_no, seller_logistic_no, bought_sku, exchange_sku, buyer_address, address, buyer_phone, buyer_logistic_name, seller_logistic_name, alipay_no, buyer_name, seller_nick
func (r *TmallExchangeGetRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r TmallExchangeGetRequest) GetFields() []string {
    return r.fields
}
