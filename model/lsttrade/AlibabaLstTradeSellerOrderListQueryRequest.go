package lsttrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单列表查看(卖家视角) API请求
alibaba.lst.trade.seller.order.list.query

卖家视角订单查询，查询授权经销商订单列表
*/
type AlibabaLstTradeSellerOrderListQueryRequest struct {
    model.Params
    // 入参
    _param   *LstTradeGetSellerOrderListParam
}

// 初始化AlibabaLstTradeSellerOrderListQueryRequest对象
func NewAlibabaLstTradeSellerOrderListQueryRequest() *AlibabaLstTradeSellerOrderListQueryRequest{
    return &AlibabaLstTradeSellerOrderListQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstTradeSellerOrderListQueryRequest) GetApiMethodName() string {
    return "alibaba.lst.trade.seller.order.list.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstTradeSellerOrderListQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参
func (r *AlibabaLstTradeSellerOrderListQueryRequest) SetParam(_param *LstTradeGetSellerOrderListParam) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaLstTradeSellerOrderListQueryRequest) GetParam() *LstTradeGetSellerOrderListParam {
    return r._param
}
