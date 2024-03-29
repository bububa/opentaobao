package icbu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductSchemaAddAPIResponse （新）商品发布新接口 API返回值
// alibaba.icbu.product.schema.add
//
// 提供发布ICBU商品的入口
type AlibabaIcbuProductSchemaAddAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuProductSchemaAddAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIcbuProductSchemaAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIcbuProductSchemaAddAPIResponseModel).Reset()
}

// AlibabaIcbuProductSchemaAddAPIResponseModel is （新）商品发布新接口 成功返回结果
type AlibabaIcbuProductSchemaAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_product_schema_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息，数组形式的字符串，用;分割，支持中英繁，按照传入的语种参数决定
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回的错误码，数组形式的字符串，用;分割
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误追踪码，请务必打印在日志中，后续排查问题请提交此错误追踪码
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 商品明文id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 调用是否成功
	BizSuccess bool `json:"biz_success,omitempty" xml:"biz_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIcbuProductSchemaAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.MsgCode = ""
	m.TraceId = ""
	m.ProductId = 0
	m.BizSuccess = false
}

var poolAlibabaIcbuProductSchemaAddAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIcbuProductSchemaAddAPIResponse)
	},
}

// GetAlibabaIcbuProductSchemaAddAPIResponse 从 sync.Pool 获取 AlibabaIcbuProductSchemaAddAPIResponse
func GetAlibabaIcbuProductSchemaAddAPIResponse() *AlibabaIcbuProductSchemaAddAPIResponse {
	return poolAlibabaIcbuProductSchemaAddAPIResponse.Get().(*AlibabaIcbuProductSchemaAddAPIResponse)
}

// ReleaseAlibabaIcbuProductSchemaAddAPIResponse 将 AlibabaIcbuProductSchemaAddAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIcbuProductSchemaAddAPIResponse(v *AlibabaIcbuProductSchemaAddAPIResponse) {
	v.Reset()
	poolAlibabaIcbuProductSchemaAddAPIResponse.Put(v)
}
