package drugtrace

import (
	"encoding/xml"

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

// AlibabaAlihealthDrugDownloadFileacceptretAPIResponseModel is 企业上传回执 成功返回结果
type AlibabaAlihealthDrugDownloadFileacceptretAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_download_fileacceptret_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *DataEntTaskResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
