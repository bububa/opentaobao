package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询P4P产品 APIResponse
alibaba.scbp.product.list

查询P4P产品
*/
type AlibabaScbpProductListAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_product_list_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 产品列表
    
    ProductList   []TopProductDto `json:"product_list,omitempty" xml:"