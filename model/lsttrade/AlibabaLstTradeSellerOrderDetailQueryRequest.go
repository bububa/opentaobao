package lsttrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单详情查看(卖家视角) API请求
alibaba.lst.trade.seller.order.detail.query

订单详情查看(卖家视角)
*/
type AlibabaLstTradeSellerOrderDetailQueryRequest struct {
    model.Params
    // 入参
    _param   *LstTradeGetSellerOrderListParam
}

// 初始化AlibabaLstTradeSellerOrderDetailQueryRequest对象
func NewAlibabaLstTradeSellerOrderDetailQueryRequest() *AlibabaLstTradeSellerOrderDetailQueryRequest{
    return &AlibabaLstTradeSellerOrderDetailQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstTradeSellerOrderDetailQueryRequest) GetApiMethodName() string {
    return "alibaba.lst.trade.seller.order.detail.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstTradeSellerOrderDetailQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参
func (r *AlibabaLstTradeSellerOrderDetailQueryRequest) SetParam(_param *LstTradeGetSellerOrderListParam) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaLstTradeSellerOrderDetailQueryRequest) GetParam() *LstTradeGetSellerOrderListParam {
    return r._param
}
