package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytgetdruginfodownloadurlAPIResponse 药品全量数据下载 API返回值
// alibaba.alihealth.drug.kyt.getdruginfo.downloadurl
//
// 药品全量数据下载
type AlibabaalihealthdrugkytgetdruginfodownloadurlAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugkytgetdruginfodownloadurlAPIResponseModel
}

// AlibabaalihealthdrugkytgetdruginfodownloadurlAPIResponseModel is 药品全量数据下载 成功返回结果
type AlibabaalihealthdrugkytgetdruginfodownloadurlAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_getdruginfo_downloadurl_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回
	Result *AlibabaalihealthdrugkytgetdruginfodownloadurlResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
