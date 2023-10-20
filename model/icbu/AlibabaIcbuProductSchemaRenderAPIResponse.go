package icbu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductSchemaRenderAPIResponse （新）获取商品信息 API返回值
// alibaba.icbu.product.schema.render
//
// 获取ICBU商品发布的字段填写规则和单个商品对应填写数据，适用于单个商品编辑场景，不包括草稿。
type AlibabaIcbuProductSchemaRenderAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuProductSchemaRenderAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIcbuProductSchemaRenderAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIcbuProductSchemaRenderAPIResponseModel).Reset()
}

// AlibabaIcbuProductSchemaRenderAPIResponseModel is （新）获取商品信息 成功返回结果
type AlibabaIcbuProductSchemaRenderAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_product_schema_render_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品发布规则和对应填写数据
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
func (m *AlibabaIcbuProductSchemaRenderAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = ""
	m.Message = ""
	m.MsgCode = ""
	m.TraceId = ""
	m.BizSuccess = false
}

var poolAlibabaIcbuProductSchemaRenderAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIcbuProductSchemaRenderAPIResponse)
	},
}

// GetAlibabaIcbuProductSchemaRenderAPIResponse 从 sync.Pool 获取 AlibabaIcbuProductSchemaRenderAPIResponse
func GetAlibabaIcbuProductSchemaRenderAPIResponse() *AlibabaIcbuProductSchemaRenderAPIResponse {
	return poolAlibabaIcbuProductSchemaRenderAPIResponse.Get().(*AlibabaIcbuProductSchemaRenderAPIResponse)
}

// ReleaseAlibabaIcbuProductSchemaRenderAPIResponse 将 AlibabaIcbuProductSchemaRenderAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIcbuProductSchemaRenderAPIResponse(v *AlibabaIcbuProductSchemaRenderAPIResponse) {
	v.Reset()
	poolAlibabaIcbuProductSchemaRenderAPIResponse.Put(v)
}
