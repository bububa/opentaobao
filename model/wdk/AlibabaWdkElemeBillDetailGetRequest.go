package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
饿了么对账单查询，带订单明细 API请求
alibaba.wdk.eleme.bill.detail.get

查询饿了么对账单信息，带订单明细
*/
type AlibabaWdkElemeBillDetailGetRequest struct {
    model.Params
    // 对账单查询参数
    _eleBillRequest   *EleBillRequest
}

// 初始化AlibabaWdkElemeBillDetailGetRequest对象
func NewAlibabaWdkElemeBillDetailGetRequest() *AlibabaWdkElemeBillDetailGetRequest{
    return &AlibabaWdkElemeBillDetailGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkElemeBillDetailGetRequest) GetApiMethodName() string {
    return "alibaba.wdk.eleme.bill.detail.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkElemeBillDetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EleBillRequest Setter
// 对账单查询参数
func (r *AlibabaWdkElemeBillDetailGetRequest) SetEleBillRequest(_eleBillRequest *EleBillRequest) error {
    r._eleBillRequest = _eleBillRequest
    r.Set("ele_bill_request", _eleBillRequest)
    return nil
}

// EleBillRequest Getter
func (r AlibabaWdkElemeBillDetailGetRequest) GetEleBillRequest() *EleBillRequest {
    return r._eleBillRequest
}
