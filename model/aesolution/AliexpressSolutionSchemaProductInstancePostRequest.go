package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.schema.product.instance.post API请求
aliexpress.solution.schema.product.instance.post

Upload product based on json schema instance.QPS(Invoke per second) for this API is limited to 100 for each appkey and 50 for each seller.
*/
type AliexpressSolutionSchemaProductInstancePostRequest struct {
    model.Params
    // Product instance data. Please note: the shipping_template_id should be replaced with your own shipping template id, which could be obtained through  https://developers.aliexpress.com/en/doc.htm?docId=43456&docType=2
    productInstanceRequest   string
}

// 初始化AliexpressSolutionSchemaProductInstancePostRequest对象
func NewAliexpressSolutionSchemaProductInstancePostRequest() *AliexpressSolutionSchemaProductInstancePostRequest{
    return &AliexpressSolutionSchemaProductInstancePostRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSolutionSchemaProductInstancePostRequest) GetApiMethodName() string {
    return "aliexpress.solution.schema.product.instance.post"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSolutionSchemaProductInstancePostRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductInstanceRequest Setter
// Product instance data. Please note: the shipping_template_id should be replaced with your own shipping template id, which could be obtained through  https://developers.aliexpress.com/en/doc.htm?docId=43456&docType=2
func (r *AliexpressSolutionSchemaProductInstancePostRequest) SetProductInstanceRequest(productInstanceRequest string) error {
    r.productInstanceRequest = productInstanceRequest
    r.Set("product_instance_request", productInstanceRequest)
    return nil
}

// ProductInstanceRequest Getter
func (r AliexpressSolutionSchemaProductInstancePostRequest) GetProductInstanceRequest() string {
    return r.productInstanceRequest
}
