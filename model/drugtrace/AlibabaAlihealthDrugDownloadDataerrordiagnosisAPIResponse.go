package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugdownloaddataerrordiagnosisAPIResponse 数据未落地诊断 API返回值
// alibaba.alihealth.drug.download.dataerrordiagnosis
//
// 阿里健康-追溯码-D2D数据未落地原因诊断
type AlibabaalihealthdrugdownloaddataerrordiagnosisAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugdownloaddataerrordiagnosisAPIResponseModel
}

// AlibabaalihealthdrugdownloaddataerrordiagnosisAPIResponseModel is 数据未落地诊断 成功返回结果
type AlibabaalihealthdrugdownloaddataerrordiagnosisAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_download_dataerrordiagnosis_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *DataEntTaskResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
