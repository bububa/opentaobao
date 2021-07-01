package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
饿了么日维度对账单查询 API请求
alibaba.wdk.eleme.bill.get

查询饿了么日维度对账单信息
*/
type AlibabaWdkElemeBillGetAPIRequest struct {
    model.Params
    // 对账单查询参数
    _eleBillRequest   *EleBillRequest
}

// 初始化AlibabaWdkElemeBillGetAPIRequest对象
func NewAlibabaWdkElemeBillGetRequest() *AlibabaWdkElemeBillGetAPIRequest{
    return &AlibabaWdkElemeBillGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkElemeBillGetAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.eleme.bill.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkElemeBillGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EleBillRequest Setter
// 对账单查询参数
func (r *AlibabaWdkElemeBillGetAPIRequest) SetEleBillRequest(_eleBillRequest *EleBillRequest) error {
    r._eleBillRequest = _eleBillRequest
    r.Set("ele_bill_request", _eleBillRequest)
    return nil
}

// EleBillRequest Getter
func (r AlibabaWdkElemeBillGetAPIRequest) GetEleBillRequest() *EleBillRequest {
    return r._eleBillRequest
}
