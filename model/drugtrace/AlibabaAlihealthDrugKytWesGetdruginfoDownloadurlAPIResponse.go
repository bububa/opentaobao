package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIResponse 药品全量数据下载 API返回值
// alibaba.alihealth.drug.kyt.wes.getdruginfo.downloadurl
//
// 药品全量数据下载
type AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIResponseModel is 药品全量数据下载 成功返回结果
type AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_wes_getdruginfo_downloadurl_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回
	Result *AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIResponse
func GetAlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIResponse() *AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIResponse {
	return poolAlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIResponse.Get().(*AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIResponse 将 AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIResponse(v *AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIResponse.Put(v)
}
