package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ascp采购价写入接口 API请求
alibaba.ascp.purchase.price.create

供应链平台采购价创建或修改接口
*/
type AlibabaAscpPurchasePriceCreateRequest struct {
    model.Params
    // 采购价创建/更新请求
    _createRequest   *Request
}

// 初始化AlibabaAscpPurchasePriceCreateRequest对象
func NewAlibabaAscpPurchasePriceCreateRequest() *AlibabaAscpPurchasePriceCreateRequest{
    return &AlibabaAscpPurchasePriceCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpPurchasePriceCreateRequest) GetApiMethodName() string {
    return "alibaba.ascp.purchase.price.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpPurchasePriceCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CreateRequest Setter
// 采购价创建/更新请求
func (r *AlibabaAscpPurchasePriceCreateRequest) SetCreateRequest(_createRequest *Request) error {
    r._createRequest = _createRequest
    r.Set("create_request", _createRequest)
    return nil
}

// CreateRequest Getter
func (r AlibabaAscpPurchasePriceCreateRequest) GetCreateRequest() *Request {
    return r._createRequest
}
