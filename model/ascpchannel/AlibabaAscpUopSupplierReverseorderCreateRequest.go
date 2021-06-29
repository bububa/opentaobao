package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家ERP发起创建销退单服务 API请求
alibaba.ascp.uop.supplier.reverseorder.create

商家在收到消费者实物退货后，在ERP发起创建销退单服务
*/
type AlibabaAscpUopSupplierReverseorderCreateRequest struct {
    model.Params
    // 逆向销退单创建请求
    reverseCreateRequest   *ReverseCreateRequest
}

// 初始化AlibabaAscpUopSupplierReverseorderCreateRequest对象
func NewAlibabaAscpUopSupplierReverseorderCreateRequest() *AlibabaAscpUopSupplierReverseorderCreateRequest{
    return &AlibabaAscpUopSupplierReverseorderCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopSupplierReverseorderCreateRequest) GetApiMethodName() string {
    return "alibaba.ascp.uop.supplier.reverseorder.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopSupplierReverseorderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReverseCreateRequest Setter
// 逆向销退单创建请求
func (r *AlibabaAscpUopSupplierReverseorderCreateRequest) SetReverseCreateRequest(reverseCreateRequest *ReverseCreateRequest) error {
    r.reverseCreateRequest = reverseCreateRequest
    r.Set("reverse_create_request", reverseCreateRequest)
    return nil
}

// ReverseCreateRequest Getter
func (r AlibabaAscpUopSupplierReverseorderCreateRequest) GetReverseCreateRequest() *ReverseCreateRequest {
    return r.reverseCreateRequest
}
