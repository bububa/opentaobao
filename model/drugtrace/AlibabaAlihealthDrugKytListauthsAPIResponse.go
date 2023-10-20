package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytlistauthsAPIResponse 企业搜索自己授权的物流企业 API返回值
// alibaba.alihealth.drug.kyt.listauths
//
// 企业搜索自己授权的物流企业
type AlibabaalihealthdrugkytlistauthsAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugkytlistauthsAPIResponseModel
}

// AlibabaalihealthdrugkytlistauthsAPIResponseModel is 企业搜索自己授权的物流企业 成功返回结果
type AlibabaalihealthdrugkytlistauthsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_listauths_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaalihealthdrugkytlistauthsResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
