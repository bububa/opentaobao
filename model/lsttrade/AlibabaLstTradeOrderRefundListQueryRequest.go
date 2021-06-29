package lsttrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询退款单列表(卖家视角) APIRequest
alibaba.lst.trade.order.refund.list.query

查询退款单列表(卖家视角)
*/
type AlibabaLstTradeOrderRefundListQueryRequest struct {
    model.Params

    // 输入参数
    param   *TopLstSupplierOrderRefundQuery 

}

func NewAlibabaLstTradeOrderRefundListQueryRequest() *AlibabaLstTradeOrderRefundListQueryRequest{
    return &AlibabaLstTradeOrderRefundListQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstTradeOrderRefundListQueryRequest) GetApiMethodName() string {
    return "alibaba.lst.trade.order.refund.list.query"
}

func (r AlibabaLstTradeOrderRefundListQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstTradeOrderRefundListQueryRequest) SetParam(param *TopLstSupplierOrderRefundQuery) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaLstTradeOrderRefundListQueryRequest) GetParam() *TopLstSupplierOrderRefundQuery {
    return r.param
}

