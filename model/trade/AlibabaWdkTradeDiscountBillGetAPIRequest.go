package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单优惠账单查询 API请求
alibaba.wdk.trade.discount.bill.get

商家查询订单优惠账单
*/
type AlibabaWdkTradeDiscountBillGetAPIRequest struct {
    model.Params
    // 请求参数
    _param0   *OrderDiscountBillQueryRequest
}

// 初始化AlibabaWdkTradeDiscountBillGetAPIRequest对象
func NewAlibabaWdkTradeDiscountBillGetRequest() *AlibabaWdkTradeDiscountBillGetAPIRequest{
    return &AlibabaWdkTradeDiscountBillGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkTradeDiscountBillGetAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.trade.discount.bill.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkTradeDiscountBillGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 请求参数
func (r *AlibabaWdkTradeDiscountBillGetAPIRequest) SetParam0(_param0 *OrderDiscountBillQueryRequest) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaWdkTradeDiscountBillGetAPIRequest) GetParam0() *OrderDiscountBillQueryRequest {
    return r._param0
}
