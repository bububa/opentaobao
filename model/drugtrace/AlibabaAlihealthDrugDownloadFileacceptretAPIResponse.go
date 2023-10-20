package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugDownloadFileacceptretAPIResponse 企业上传回执 API返回值
// alibaba.alihealth.drug.download.fileacceptret
//
// 拿到企业下载回执，将企业已下载的和未下载成功的条目都相应的改变状态
type AlibabaAlihealthDrugDownloadFileacceptretAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugDownloadFileacceptretAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugDownloadFileacceptretAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugDownloadFileacceptretAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugDownloadFileacceptretAPIResponseModel is 企业上传回执 成功返回结果
type AlibabaAlihealthDrugDownloadFileacceptretAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_download_fileacceptret_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *DataEntTaskResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugDownloadFileacceptretAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugDownloadFileacceptretAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugDownloadFileacceptretAPIResponse)
	},
}

// GetAlibabaAlihealthDrugDownloadFileacceptretAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugDownloadFileacceptretAPIResponse
func GetAlibabaAlihealthDrugDownloadFileacceptretAPIResponse() *AlibabaAlihealthDrugDownloadFileacceptretAPIResponse {
	return poolAlibabaAlihealthDrugDownloadFileacceptretAPIResponse.Get().(*AlibabaAlihealthDrugDownloadFileacceptretAPIResponse)
}

// ReleaseAlibabaAlihealthDrugDownloadFileacceptretAPIResponse 将 AlibabaAlihealthDrugDownloadFileacceptretAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugDownloadFileacceptretAPIResponse(v *AlibabaAlihealthDrugDownloadFileacceptretAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugDownloadFileacceptretAPIResponse.Put(v)
}
