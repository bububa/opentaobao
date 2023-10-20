package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIResponse 获取重点追溯品种明细下载URL API返回值
// alibaba.alihealth.drugtrace.top.lsyd.getkeyflagdruginfo.downloadurl
//
// 获取重点追溯品种明细下载URL
type AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIResponseModel is 获取重点追溯品种明细下载URL 成功返回结果
type AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_lsyd_getkeyflagdruginfo_downloadurl_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回
	Result *AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIResponse)
	},
}

// GetAlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIResponse
func GetAlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIResponse() *AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIResponse {
	return poolAlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIResponse.Get().(*AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIResponse)
}

// ReleaseAlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIResponse 将 AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIResponse(v *AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIResponse.Put(v)
}
