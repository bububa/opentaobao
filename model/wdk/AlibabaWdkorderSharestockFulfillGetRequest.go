package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商户订单履约数据获取 API请求
alibaba.wdkorder.sharestock.fulfill.get

商户订单履约数据获取
*/
type AlibabaWdkorderSharestockFulfillGetRequest struct {
    model.Params
    // 履约单ID
    fulfillOrderId   string
}

// 初始化AlibabaWdkorderSharestockFulfillGetRequest对象
func NewAlibabaWdkorderSharestockFulfillGetRequest() *AlibabaWdkorderSharestockFulfillGetRequest{
    return &AlibabaWdkorderSharestockFulfillGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkorderSharestockFulfillGetRequest) GetApiMethodName() string {
    return "alibaba.wdkorder.sharestock.fulfill.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkorderSharestockFulfillGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FulfillOrderId Setter
// 履约单ID
func (r *AlibabaWdkorderSharestockFulfillGetRequest) SetFulfillOrderId(fulfillOrderId string) error {
    r.fulfillOrderId = fulfillOrderId
    r.Set("fulfill_order_id", fulfillOrderId)
    return nil
}

// FulfillOrderId Getter
func (r AlibabaWdkorderSharestockFulfillGetRequest) GetFulfillOrderId() string {
    return r.fulfillOrderId
}
