package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家ERP发起创建销退单服务 APIRequest
alibaba.ascp.uop.supplier.reverseorder.create

商家在收到消费者实物退货后，在ERP发起创建销退单服务
*/
type AlibabaAscpUopSupplierReverseorderCreateRequest struct {
    model.Params

    // 逆向销退单创建请求
    reverseCreateRequest   *ReverseCreateRequest 

}

func NewAlibabaAscpUopSupplierReverseorderCreateRequest() *AlibabaAscpUopSupplierReverseorderCreateRequest{
    return &AlibabaAscpUopSupplierReverseorderCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpUopSupplierReverseorderCreateRequest) GetApiMethodName() string {
    return "alibaba.ascp.uop.supplier.reverseorder.create"
}

func (r AlibabaAscpUopSupplierReverseorderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpUopSupplierReverseorderCreateRequest) SetReverseCreateRequest(reverseCreateRequest *ReverseCreateRequest) error {
    r.reverseCreateRequest = reverseCreateRequest
    r.Set("reverse_create_request", reverseCreateRequest)
    return nil
}

func (r AlibabaAscpUopSupplierReverseorderCreateRequest) GetReverseCreateRequest() *ReverseCreateRequest {
    return r.reverseCreateRequest
}

