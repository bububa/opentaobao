package drugtrace

import (
	"encoding/xml"

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

// AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIResponseModel is 药品全量数据下载 成功返回结果
type AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_wes_getdruginfo_downloadurl_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回
	Result *AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
