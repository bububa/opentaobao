package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIResponse 数据未落地诊断 API返回值
// alibaba.alihealth.drug.download.dataerrordiagnosis
//
// 阿里健康-追溯码-D2D数据未落地原因诊断
type AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIResponseModel is 数据未落地诊断 成功返回结果
type AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_download_dataerrordiagnosis_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *DataEntTaskResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugDownloadDataerrordiagnosisAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIResponse)
	},
}

// GetAlibabaAlihealthDrugDownloadDataerrordiagnosisAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIResponse
func GetAlibabaAlihealthDrugDownloadDataerrordiagnosisAPIResponse() *AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIResponse {
	return poolAlibabaAlihealthDrugDownloadDataerrordiagnosisAPIResponse.Get().(*AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIResponse)
}

// ReleaseAlibabaAlihealthDrugDownloadDataerrordiagnosisAPIResponse 将 AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugDownloadDataerrordiagnosisAPIResponse(v *AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugDownloadDataerrordiagnosisAPIResponse.Put(v)
}
