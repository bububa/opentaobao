package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.schema.product.instance.post APIRequest
aliexpress.solution.schema.product.instance.post

Upload product based on json schema instance.QPS(Invoke per second) for this API is limited to 100 for each appkey and 50 for each seller.
*/
type AliexpressSolutionSchemaProductInstancePostRequest struct {
    model.Params

    // Product instance data. Please note: the shipping_template_id should be replaced with your own shipping template id, which could be obtained through  https://developers.aliexpress.com/en/doc.htm?docId=43456&docType=2
    productInstanceRequest   string 

}

func NewAliexpressSolutionSchemaProductInstancePostRequest() *AliexpressSolutionSchemaProductInstancePostRequest{
    return &AliexpressSolutionSchemaProductInstancePostRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressSolutionSchemaProductInstancePostRequest) GetApiMethodName() string {
    return "aliexpress.solution.schema.product.instance.post"
}

func (r AliexpressSolutionSchemaProductInstancePostRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressSolutionSchemaProductInstancePostRequest) SetProductInstanceRequest(productInstanceRequest string) error {
    r.productInstanceRequest = productInstanceRequest
    r.Set("product_instance_request", productInstanceRequest)
    return nil
}

func (r AliexpressSolutionSchemaProductInstancePostRequest) GetProductInstanceRequest() string {
    return r.productInstanceRequest
}

