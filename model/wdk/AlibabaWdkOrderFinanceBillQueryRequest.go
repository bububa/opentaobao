package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
资金合规商家账单 API请求
alibaba.wdk.order.finance.bill.query

拉取资金合规商家账单
*/
type AlibabaWdkOrderFinanceBillQueryRequest struct {
    model.Params
    // 入参
    _billQueryRequest   *WdkOpenOrderFinanceBillQueryRequest
}

// 初始化AlibabaWdkOrderFinanceBillQueryRequest对象
func NewAlibabaWdkOrderFinanceBillQueryRequest() *AlibabaWdkOrderFinanceBillQueryRequest{
    return &AlibabaWdkOrderFinanceBillQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkOrderFinanceBillQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.order.finance.bill.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkOrderFinanceBillQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BillQueryRequest Setter
// 入参
func (r *AlibabaWdkOrderFinanceBillQueryRequest) SetBillQueryRequest(_billQueryRequest *WdkOpenOrderFinanceBillQueryRequest) error {
    r._billQueryRequest = _billQueryRequest
    r.Set("bill_query_request", _billQueryRequest)
    return nil
}

// BillQueryRequest Getter
func (r AlibabaWdkOrderFinanceBillQueryRequest) GetBillQueryRequest() *WdkOpenOrderFinanceBillQueryRequest {
    return r._billQueryRequest
}
