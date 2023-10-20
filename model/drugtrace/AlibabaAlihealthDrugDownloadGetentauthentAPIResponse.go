package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugDownloadGetentauthentAPIResponse 获取授权企业列表 API返回值
// alibaba.alihealth.drug.download.getentauthent
//
// D2D数据落地获取授权企业列表
type AlibabaAlihealthDrugDownloadGetentauthentAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugDownloadGetentauthentAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugDownloadGetentauthentAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugDownloadGetentauthentAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugDownloadGetentauthentAPIResponseModel is 获取授权企业列表 成功返回结果
type AlibabaAlihealthDrugDownloadGetentauthentAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_download_getentauthent_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihealthDrugDownloadGetentauthentResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugDownloadGetentauthentAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugDownloadGetentauthentAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugDownloadGetentauthentAPIResponse)
	},
}

// GetAlibabaAlihealthDrugDownloadGetentauthentAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugDownloadGetentauthentAPIResponse
func GetAlibabaAlihealthDrugDownloadGetentauthentAPIResponse() *AlibabaAlihealthDrugDownloadGetentauthentAPIResponse {
	return poolAlibabaAlihealthDrugDownloadGetentauthentAPIResponse.Get().(*AlibabaAlihealthDrugDownloadGetentauthentAPIResponse)
}

// ReleaseAlibabaAlihealthDrugDownloadGetentauthentAPIResponse 将 AlibabaAlihealthDrugDownloadGetentauthentAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugDownloadGetentauthentAPIResponse(v *AlibabaAlihealthDrugDownloadGetentauthentAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugDownloadGetentauthentAPIResponse.Put(v)
}
