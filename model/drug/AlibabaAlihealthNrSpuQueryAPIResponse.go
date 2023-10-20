package drug

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthNrSpuQueryAPIResponse 获取标品库标品信息 API返回值
// alibaba.alihealth.nr.spu.query
//
// 提供给ERP使用的，获取健康标品库信息
type AlibabaAlihealthNrSpuQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthNrSpuQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthNrSpuQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthNrSpuQueryAPIResponseModel).Reset()
}

// AlibabaAlihealthNrSpuQueryAPIResponseModel is 获取标品库标品信息 成功返回结果
type AlibabaAlihealthNrSpuQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_nr_spu_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *ResponseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthNrSpuQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthNrSpuQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthNrSpuQueryAPIResponse)
	},
}

// GetAlibabaAlihealthNrSpuQueryAPIResponse 从 sync.Pool 获取 AlibabaAlihealthNrSpuQueryAPIResponse
func GetAlibabaAlihealthNrSpuQueryAPIResponse() *AlibabaAlihealthNrSpuQueryAPIResponse {
	return poolAlibabaAlihealthNrSpuQueryAPIResponse.Get().(*AlibabaAlihealthNrSpuQueryAPIResponse)
}

// ReleaseAlibabaAlihealthNrSpuQueryAPIResponse 将 AlibabaAlihealthNrSpuQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthNrSpuQueryAPIResponse(v *AlibabaAlihealthNrSpuQueryAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthNrSpuQueryAPIResponse.Put(v)
}
