package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIResponse 药品全量数据下载 API返回值
// alibaba.alihealth.drug.kyt.getdruginfo.downloadurl
//
// 药品全量数据下载
type AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIResponseModel is 药品全量数据下载 成功返回结果
type AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_getdruginfo_downloadurl_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回
	Result *AlibabaAlihealthDrugKytGetdruginfoDownloadurlResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIResponse
func GetAlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIResponse() *AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIResponse {
	return poolAlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIResponse.Get().(*AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIResponse 将 AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIResponse(v *AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIResponse.Put(v)
}
