package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.schema.product.full.update API请求
aliexpress.solution.schema.product.full.update

Schema interface for product full update. QPS(Invoke per second) for this API is limited to 100 for each appkey and 50 for each seller.
*/
type AliexpressSolutionSchemaProductFullUpdateRequest struct {
    model.Params
    // Product full update request. To learn how to generate the content, please refer to https://developers.aliexpress.com/en/doc.htm?docId=109760&docType=1.  Be aware that the aliexpress_product_id field should be replaced by the product ID belonged to the seller.
    schemaFullUpdateRequest   string
}

// 初始化AliexpressSolutionSchemaProductFullUpdateRequest对象
func NewAliexpressSolutionSchemaProductFullUpdateRequest() *AliexpressSolutionSchemaProductFullUpdateRequest{
    return &AliexpressSolutionSchemaProductFullUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSolutionSchemaProductFullUpdateRequest) GetApiMethodName() string {
    return "aliexpress.solution.schema.product.full.update"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSolutionSchemaProductFullUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SchemaFullUpdateRequest Setter
// Product full update request. To learn how to generate the content, please refer to https://developers.aliexpress.com/en/doc.htm?docId=109760&docType=1.  Be aware that the aliexpress_product_id field should be replaced by the product ID belonged to the seller.
func (r *AliexpressSolutionSchemaProductFullUpdateRequest) SetSchemaFullUpdateRequest(schemaFullUpdateRequest string) error {
    r.schemaFullUpdateRequest = schemaFullUpdateRequest
    r.Set("schema_full_update_request", schemaFullUpdateRequest)
    return nil
}

// SchemaFullUpdateRequest Getter
func (r AliexpressSolutionSchemaProductFullUpdateRequest) GetSchemaFullUpdateRequest() string {
    return r.schemaFullUpdateRequest
}
