package icbu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWholesaleShippinglineTemplateListAPIResponse 获取运费模板 API返回值
// alibaba.wholesale.shippingline.template.list
//
// 查询运费模板信息
type AlibabaWholesaleShippinglineTemplateListAPIResponse struct {
	model.CommonResponse
	AlibabaWholesaleShippinglineTemplateListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWholesaleShippinglineTemplateListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWholesaleShippinglineTemplateListAPIResponseModel).Reset()
}

// AlibabaWholesaleShippinglineTemplateListAPIResponseModel is 获取运费模板 成功返回结果
type AlibabaWholesaleShippinglineTemplateListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wholesale_shippingline_template_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 运费模板列表
	ListTemplateResponse *ListTemplateAPIResult `json:"list_template_response,omitempty" xml:"list_template_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWholesaleShippinglineTemplateListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ListTemplateResponse = nil
}

var poolAlibabaWholesaleShippinglineTemplateListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWholesaleShippinglineTemplateListAPIResponse)
	},
}

// GetAlibabaWholesaleShippinglineTemplateListAPIResponse 从 sync.Pool 获取 AlibabaWholesaleShippinglineTemplateListAPIResponse
func GetAlibabaWholesaleShippinglineTemplateListAPIResponse() *AlibabaWholesaleShippinglineTemplateListAPIResponse {
	return poolAlibabaWholesaleShippinglineTemplateListAPIResponse.Get().(*AlibabaWholesaleShippinglineTemplateListAPIResponse)
}

// ReleaseAlibabaWholesaleShippinglineTemplateListAPIResponse 将 AlibabaWholesaleShippinglineTemplateListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWholesaleShippinglineTemplateListAPIResponse(v *AlibabaWholesaleShippinglineTemplateListAPIResponse) {
	v.Reset()
	poolAlibabaWholesaleShippinglineTemplateListAPIResponse.Put(v)
}
