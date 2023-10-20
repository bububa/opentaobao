package icbu

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbuproductschemagetAPIResponse （新）ICBU商品发布schema接口 API返回值
// alibaba.icbu.product.schema.get
//
// 获取ICBU商品发布的页面规则和填写字段，适用于新发商品
type AlibabaicbuproductschemagetAPIResponse struct {
	model.CommonResponse
	AlibabaicbuproductschemagetAPIResponseModel
}

// AlibabaicbuproductschemagetAPIResponseModel is （新）ICBU商品发布schema接口 成功返回结果
type AlibabaicbuproductschemagetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_product_schema_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品发布规则
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息，数组形式的字符串，用;分割，支持中英繁，按照传入的语种参数决定
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回的错误码，数组形式的字符串，用;分割
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误追踪码，请务必打印在日志中，后续排查问题请提交此错误追踪码
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 请求是否成功
	BizSuccess bool `json:"biz_success,omitempty" xml:"biz_success,omitempty"`
}
