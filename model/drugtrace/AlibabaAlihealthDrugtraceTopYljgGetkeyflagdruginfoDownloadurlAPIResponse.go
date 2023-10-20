package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIResponse 获取重点追溯品种明细下载URL API返回值
// alibaba.alihealth.drugtrace.top.yljg.getkeyflagdruginfo.downloadurl
//
// 获取重点追溯品种明细下载URL
type AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIResponseModel is 获取重点追溯品种明细下载URL 成功返回结果
type AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_yljg_getkeyflagdruginfo_downloadurl_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回
	Result *AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIResponse)
	},
}

// GetAlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIResponse
func GetAlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIResponse() *AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIResponse {
	return poolAlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIResponse.Get().(*AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIResponse)
}

// ReleaseAlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIResponse 将 AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIResponse(v *AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIResponse.Put(v)
}
