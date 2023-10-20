package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytScqyUploadrelationAPIResponse 关联关系上传 API返回值
// alibaba.alihealth.drug.kyt.scqy.uploadrelation
//
// 关联关系上传
type AlibabaAlihealthDrugKytScqyUploadrelationAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytScqyUploadrelationAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytScqyUploadrelationAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytScqyUploadrelationAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytScqyUploadrelationAPIResponseModel is 关联关系上传 成功返回结果
type AlibabaAlihealthDrugKytScqyUploadrelationAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_scqy_uploadrelation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihealthDrugKytScqyUploadrelationResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytScqyUploadrelationAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytScqyUploadrelationAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytScqyUploadrelationAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytScqyUploadrelationAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytScqyUploadrelationAPIResponse
func GetAlibabaAlihealthDrugKytScqyUploadrelationAPIResponse() *AlibabaAlihealthDrugKytScqyUploadrelationAPIResponse {
	return poolAlibabaAlihealthDrugKytScqyUploadrelationAPIResponse.Get().(*AlibabaAlihealthDrugKytScqyUploadrelationAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytScqyUploadrelationAPIResponse 将 AlibabaAlihealthDrugKytScqyUploadrelationAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytScqyUploadrelationAPIResponse(v *AlibabaAlihealthDrugKytScqyUploadrelationAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytScqyUploadrelationAPIResponse.Put(v)
}
