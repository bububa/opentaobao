package icbu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品按字段更新 APIResponse
alibaba.icbu.product.update.field

按字段修改国际站商品，支持询盘商品和在线批发商品，支持英文商品和多语言商品
*/
type AlibabaIcbuProductUpdateFieldAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_icbu_product_update_field_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 加密后的产品ID
    
    ProductId   string `json:"product_id,omitempty" xml:"