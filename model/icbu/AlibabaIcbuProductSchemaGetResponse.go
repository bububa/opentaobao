package icbu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
（新）ICBU商品发布schema接口 APIResponse
alibaba.icbu.product.schema.get

获取ICBU商品发布的页面规则和填写字段，适用于新发商品
*/
type AlibabaIcbuProductSchemaGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaIcbuProductSchemaGetResponse `json:"alibaba_icbu_product_schema_get_response,omitempty"`
}

type AlibabaIcbuProductSchemaGetResponse struct {

    // 商品发布规则
    Data   string `json:"data,omitempty"`

    // 错误信息，数组形式的字符串，用;分割，支持中英繁，按照传入的语种参数决定
    Message   string `json:"message,omitempty"`

    // 返回的错误码，数组形式的字符串，用;分割
    MsgCode   string `json:"msg_code,omitempty"`

    // 请求是否成功
    BizSuccess   bool `json:"biz_success,omitempty"`

    // 错误追踪码，请务必打印在日志中，后续排查问题请提交此错误追踪码
    TraceId   string `json:"trace_id,omitempty"`

}
