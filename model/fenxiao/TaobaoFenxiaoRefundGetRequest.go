package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询采购单退款信息 API请求
taobao.fenxiao.refund.get

分销商或供应商可以查询某子单的退款信息，以及下游订单的退款信息
*/
type TaobaoFenxiaoRefundGetRequest struct {
    model.Params
    // 要查询的退款子单的id
    subOrderId   int64
    // 是否查询下游买家的退款信息
    querySellerRefund   bool
}

// 初始化TaobaoFenxiaoRefundGetRequest对象
func NewTaobaoFenxiaoRefundGetRequest() *TaobaoFenxiaoRefundGetRequest{
    return &TaobaoFenxiaoRefundGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoRefundGetRequest) GetApiMethodName() string {
    return "taobao.fenxiao.refund.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoRefundGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SubOrderId Setter
// 要查询的退款子单的id
func (r *TaobaoFenxiaoRefundGetRequest) SetSubOrderId(subOrderId int64) error {
    r.subOrderId = subOrderId
    r.Set("sub_order_id", subOrderId)
    return nil
}

// SubOrderId Getter
func (r TaobaoFenxiaoRefundGetRequest) GetSubOrderId() int64 {
    return r.subOrderId
}
// QuerySellerRefund Setter
// 是否查询下游买家的退款信息
func (r *TaobaoFenxiaoRefundGetRequest) SetQuerySellerRefund(querySellerRefund bool) error {
    r.querySellerRefund = querySellerRefund
    r.Set("query_seller_refund", querySellerRefund)
    return nil
}

// QuerySellerRefund Getter
func (r TaobaoFenxiaoRefundGetRequest) GetQuerySellerRefund() bool {
    return r.querySellerRefund
}
