package icbu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）商品发布增量更新接口 APIResponse
alibaba.icbu.product.schema.update

商品更新接口，方式为增量更新，只更新传入的字段
*/
type AlibabaIcbuProductSchemaUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaIcbuProductSchemaUpdateResponse
}

type AlibabaIcbuProductSchemaUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_icbu_product_schema_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品明文id
    
    ProductId   int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`

    
    // 错误信息，数组形式的字符串，用;分割，支持中英繁，按照传入的语种参数决定
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 返回的错误码，数组形式的字符串，用;分割
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
    // 调用是否成功
    
    BizSuccess   bool `json:"biz_success,omitempty" xml:"biz_success,omitempty"`

    
    // 错误追踪码，请务必打印在日志中，后续排查问题请提交此错误追踪码
    
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`

    
}
