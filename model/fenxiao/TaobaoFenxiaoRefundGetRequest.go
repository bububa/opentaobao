package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询采购单退款信息 APIRequest
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

func NewTaobaoFenxiaoRefundGetRequest() *TaobaoFenxiaoRefundGetRequest{
    return &TaobaoFenxiaoRefundGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFenxiaoRefundGetRequest) GetApiMethodName() string {
    return "taobao.fenxiao.refund.get"
}

func (r TaobaoFenxiaoRefundGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFenxiaoRefundGetRequest) SetSubOrderId(subOrderId int64) error {
    r.subOrderId = subOrderId
    r.Set("sub_order_id", subOrderId)
    return nil
}

func (r TaobaoFenxiaoRefundGetRequest) GetSubOrderId() int64 {
    return r.subOrderId
}

func (r *TaobaoFenxiaoRefundGetRequest) SetQuerySellerRefund(querySellerRefund bool) error {
    r.querySellerRefund = querySellerRefund
    r.Set("query_seller_refund", querySellerRefund)
    return nil
}

func (r TaobaoFenxiaoRefundGetRequest) GetQuerySellerRefund() bool {
    return r.querySellerRefund
}

