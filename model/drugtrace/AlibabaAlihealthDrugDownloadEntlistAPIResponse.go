package drugtrace

import (
	"encoding/xml"

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

// AlibabaAlihealthDrugDownloadEntlistAPIResponseModel is 企业下载列表 成功返回结果
type AlibabaAlihealthDrugDownloadEntlistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_download_entlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAlihealthDrugDownloadEntlistResult `json:"result,omitempty" xml:"result,omitempty"`
}
