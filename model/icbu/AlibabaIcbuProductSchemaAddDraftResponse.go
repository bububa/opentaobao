package icbu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
（新）ICBU商品发布草稿 APIResponse
alibaba.icbu.product.schema.add.draft

提供发布ICBU商品草稿的入口
*/
type AlibabaIcbuProductSchemaAddDraftAPIResponse struct {
    model.CommonResponse
    Response *AlibabaIcbuProductSchemaAddDraftResponse `json:"alibaba_icbu_product_schema_add_draft_response,omitempty"`
}

type AlibabaIcbuProductSchemaAddDraftResponse struct {

    // 商品草稿明文id
    ProductId   int64 `json:"product_id,omitempty"`

    // 错误信息，数组形式的字符串，用;分割，支持中英繁，按照传入的语种参数决定
    Message   string `json:"message,omitempty"`

    // 返回的错误码，数组形式的字符串，用;分割
    MsgCode   string `json:"msg_code,omitempty"`

    // 调用是否成功
    BizSuccess   bool `json:"biz_success,omitempty"`

    // 错误追踪码，请务必打印在日志中，后续排查问题请提交此错误追踪码
    TraceId   string `json:"trace_id,omitempty"`

}
