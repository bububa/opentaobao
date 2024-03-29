package icbu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductSchemaGetAPIResponse （新）ICBU商品发布schema接口 API返回值
// alibaba.icbu.product.schema.get
//
// 获取ICBU商品发布的页面规则和填写字段，适用于新发商品
type AlibabaIcbuProductSchemaGetAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuProductSchemaGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIcbuProductSchemaGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIcbuProductSchemaGetAPIResponseModel).Reset()
}

// AlibabaIcbuProductSchemaGetAPIResponseModel is （新）ICBU商品发布schema接口 成功返回结果
type AlibabaIcbuProductSchemaGetAPIResponseModel struct {
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

// Reset 清空结构体
func (m *AlibabaIcbuProductSchemaGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = ""
	m.Message = ""
	m.MsgCode = ""
	m.TraceId = ""
	m.BizSuccess = false
}

var poolAlibabaIcbuProductSchemaGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIcbuProductSchemaGetAPIResponse)
	},
}

// GetAlibabaIcbuProductSchemaGetAPIResponse 从 sync.Pool 获取 AlibabaIcbuProductSchemaGetAPIResponse
func GetAlibabaIcbuProductSchemaGetAPIResponse() *AlibabaIcbuProductSchemaGetAPIResponse {
	return poolAlibabaIcbuProductSchemaGetAPIResponse.Get().(*AlibabaIcbuProductSchemaGetAPIResponse)
}

// ReleaseAlibabaIcbuProductSchemaGetAPIResponse 将 AlibabaIcbuProductSchemaGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIcbuProductSchemaGetAPIResponse(v *AlibabaIcbuProductSchemaGetAPIResponse) {
	v.Reset()
	poolAlibabaIcbuProductSchemaGetAPIResponse.Put(v)
}
