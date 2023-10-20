package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugtracetoplsydgetkeyflagdruginfodownloadurlAPIResponse 获取重点追溯品种明细下载URL API返回值
// alibaba.alihealth.drugtrace.top.lsyd.getkeyflagdruginfo.downloadurl
//
// 获取重点追溯品种明细下载URL
type AlibabaalihealthdrugtracetoplsydgetkeyflagdruginfodownloadurlAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugtracetoplsydgetkeyflagdruginfodownloadurlAPIResponseModel
}

// AlibabaalihealthdrugtracetoplsydgetkeyflagdruginfodownloadurlAPIResponseModel is 获取重点追溯品种明细下载URL 成功返回结果
type AlibabaalihealthdrugtracetoplsydgetkeyflagdruginfodownloadurlAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_lsyd_getkeyflagdruginfo_downloadurl_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回
	Result *AlibabaalihealthdrugtracetoplsydgetkeyflagdruginfodownloadurlResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
