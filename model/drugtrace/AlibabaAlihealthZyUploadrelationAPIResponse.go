package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthZyUploadrelationAPIResponse 中药片关联关系上传 API返回值
// alibaba.alihealth.zy.uploadrelation
//
// 中药片关联关系上传
type AlibabaAlihealthZyUploadrelationAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthZyUploadrelationAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthZyUploadrelationAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthZyUploadrelationAPIResponseModel).Reset()
}

// AlibabaAlihealthZyUploadrelationAPIResponseModel is 中药片关联关系上传 成功返回结果
type AlibabaAlihealthZyUploadrelationAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_zy_uploadrelation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihealthZyUploadrelationResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthZyUploadrelationAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthZyUploadrelationAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthZyUploadrelationAPIResponse)
	},
}

// GetAlibabaAlihealthZyUploadrelationAPIResponse 从 sync.Pool 获取 AlibabaAlihealthZyUploadrelationAPIResponse
func GetAlibabaAlihealthZyUploadrelationAPIResponse() *AlibabaAlihealthZyUploadrelationAPIResponse {
	return poolAlibabaAlihealthZyUploadrelationAPIResponse.Get().(*AlibabaAlihealthZyUploadrelationAPIResponse)
}

// ReleaseAlibabaAlihealthZyUploadrelationAPIResponse 将 AlibabaAlihealthZyUploadrelationAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthZyUploadrelationAPIResponse(v *AlibabaAlihealthZyUploadrelationAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthZyUploadrelationAPIResponse.Put(v)
}
