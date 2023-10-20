package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugDownloadEntlistAPIResponse 企业下载列表 API返回值
// alibaba.alihealth.drug.download.entlist
//
// 获取企业的下载文件列表
type AlibabaAlihealthDrugDownloadEntlistAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugDownloadEntlistAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugDownloadEntlistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugDownloadEntlistAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugDownloadEntlistAPIResponseModel is 企业下载列表 成功返回结果
type AlibabaAlihealthDrugDownloadEntlistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_download_entlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAlihealthDrugDownloadEntlistResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugDownloadEntlistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugDownloadEntlistAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugDownloadEntlistAPIResponse)
	},
}

// GetAlibabaAlihealthDrugDownloadEntlistAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugDownloadEntlistAPIResponse
func GetAlibabaAlihealthDrugDownloadEntlistAPIResponse() *AlibabaAlihealthDrugDownloadEntlistAPIResponse {
	return poolAlibabaAlihealthDrugDownloadEntlistAPIResponse.Get().(*AlibabaAlihealthDrugDownloadEntlistAPIResponse)
}

// ReleaseAlibabaAlihealthDrugDownloadEntlistAPIResponse 将 AlibabaAlihealthDrugDownloadEntlistAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugDownloadEntlistAPIResponse(v *AlibabaAlihealthDrugDownloadEntlistAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugDownloadEntlistAPIResponse.Put(v)
}
