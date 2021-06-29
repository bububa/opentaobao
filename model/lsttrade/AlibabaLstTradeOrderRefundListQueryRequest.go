package lsttrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询退款单列表(卖家视角) API请求
alibaba.lst.trade.order.refund.list.query

查询退款单列表(卖家视角)
*/
type AlibabaLstTradeOrderRefundListQueryRequest struct {
    model.Params
    // 输入参数
    _param   *TopLstSupplierOrderRefundQuery
}

// 初始化AlibabaLstTradeOrderRefundListQueryRequest对象
func NewAlibabaLstTradeOrderRefundListQueryRequest() *AlibabaLstTradeOrderRefundListQueryRequest{
    return &AlibabaLstTradeOrderRefundListQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstTradeOrderRefundListQueryRequest) GetApiMethodName() string {
    return "alibaba.lst.trade.order.refund.list.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstTradeOrderRefundListQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 输入参数
func (r *AlibabaLstTradeOrderRefundListQueryRequest) SetParam(_param *TopLstSupplierOrderRefundQuery) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaLstTradeOrderRefundListQueryRequest) GetParam() *TopLstSupplierOrderRefundQuery {
    return r._param
}
