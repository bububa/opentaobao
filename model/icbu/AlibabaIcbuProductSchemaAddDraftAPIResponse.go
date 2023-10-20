package icbu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductSchemaAddDraftAPIResponse （新）ICBU商品发布草稿 API返回值
// alibaba.icbu.product.schema.add.draft
//
// 提供发布ICBU商品草稿的入口
type AlibabaIcbuProductSchemaAddDraftAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuProductSchemaAddDraftAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIcbuProductSchemaAddDraftAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIcbuProductSchemaAddDraftAPIResponseModel).Reset()
}

// AlibabaIcbuProductSchemaAddDraftAPIResponseModel is （新）ICBU商品发布草稿 成功返回结果
type AlibabaIcbuProductSchemaAddDraftAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_product_schema_add_draft_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息，数组形式的字符串，用;分割，支持中英繁，按照传入的语种参数决定
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回的错误码，数组形式的字符串，用;分割
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误追踪码，请务必打印在日志中，后续排查问题请提交此错误追踪码
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 商品草稿明文id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 调用是否成功
	BizSuccess bool `json:"biz_success,omitempty" xml:"biz_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIcbuProductSchemaAddDraftAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.MsgCode = ""
	m.TraceId = ""
	m.ProductId = 0
	m.BizSuccess = false
}

var poolAlibabaIcbuProductSchemaAddDraftAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIcbuProductSchemaAddDraftAPIResponse)
	},
}

// GetAlibabaIcbuProductSchemaAddDraftAPIResponse 从 sync.Pool 获取 AlibabaIcbuProductSchemaAddDraftAPIResponse
func GetAlibabaIcbuProductSchemaAddDraftAPIResponse() *AlibabaIcbuProductSchemaAddDraftAPIResponse {
	return poolAlibabaIcbuProductSchemaAddDraftAPIResponse.Get().(*AlibabaIcbuProductSchemaAddDraftAPIResponse)
}

// ReleaseAlibabaIcbuProductSchemaAddDraftAPIResponse 将 AlibabaIcbuProductSchemaAddDraftAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIcbuProductSchemaAddDraftAPIResponse(v *AlibabaIcbuProductSchemaAddDraftAPIResponse) {
	v.Reset()
	poolAlibabaIcbuProductSchemaAddDraftAPIResponse.Put(v)
}
