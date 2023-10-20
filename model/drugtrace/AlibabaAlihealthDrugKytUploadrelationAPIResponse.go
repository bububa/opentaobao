package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytUploadrelationAPIResponse 关联关系上传 API返回值
// alibaba.alihealth.drug.kyt.uploadrelation
//
// 关联关系上传
type AlibabaAlihealthDrugKytUploadrelationAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytUploadrelationAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytUploadrelationAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytUploadrelationAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytUploadrelationAPIResponseModel is 关联关系上传 成功返回结果
type AlibabaAlihealthDrugKytUploadrelationAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_uploadrelation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihealthDrugKytUploadrelationResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytUploadrelationAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytUploadrelationAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytUploadrelationAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytUploadrelationAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytUploadrelationAPIResponse
func GetAlibabaAlihealthDrugKytUploadrelationAPIResponse() *AlibabaAlihealthDrugKytUploadrelationAPIResponse {
	return poolAlibabaAlihealthDrugKytUploadrelationAPIResponse.Get().(*AlibabaAlihealthDrugKytUploadrelationAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytUploadrelationAPIResponse 将 AlibabaAlihealthDrugKytUploadrelationAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytUploadrelationAPIResponse(v *AlibabaAlihealthDrugKytUploadrelationAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytUploadrelationAPIResponse.Put(v)
}
