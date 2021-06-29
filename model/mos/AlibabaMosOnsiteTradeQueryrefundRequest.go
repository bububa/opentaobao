package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退款查询 API请求
alibaba.mos.onsite.trade.queryrefund

商户可使用该接口查询退款请求是否执行成功。
*/
type AlibabaMosOnsiteTradeQueryrefundRequest struct {
    model.Params
    // 退款外部流水号
    _outRequestNo   string
    // 订单号。可能为外部订单号，也可能为喵街订单号
    _orderNo   string
}

// 初始化AlibabaMosOnsiteTradeQueryrefundRequest对象
func NewAlibabaMosOnsiteTradeQueryrefundRequest() *AlibabaMosOnsiteTradeQueryrefundRequest{
    return &AlibabaMosOnsiteTradeQueryrefundRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMosOnsiteTradeQueryrefundRequest) GetApiMethodName() string {
    return "alibaba.mos.onsite.trade.queryrefund"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMosOnsiteTradeQueryrefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutRequestNo Setter
// 退款外部流水号
func (r *AlibabaMosOnsiteTradeQueryrefundRequest) SetOutRequestNo(_outRequestNo string) error {
    r._outRequestNo = _outRequestNo
    r.Set("out_request_no", _outRequestNo)
    return nil
}

// OutRequestNo Getter
func (r AlibabaMosOnsiteTradeQueryrefundRequest) GetOutRequestNo() string {
    return r._outRequestNo
}
// OrderNo Setter
// 订单号。可能为外部订单号，也可能为喵街订单号
func (r *AlibabaMosOnsiteTradeQueryrefundRequest) SetOrderNo(_orderNo string) error {
    r._orderNo = _orderNo
    r.Set("order_no", _orderNo)
    return nil
}

// OrderNo Getter
func (r AlibabaMosOnsiteTradeQueryrefundRequest) GetOrderNo() string {
    return r._orderNo
}
