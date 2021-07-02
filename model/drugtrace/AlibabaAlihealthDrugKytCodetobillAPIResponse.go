package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytCodetobillAPIResponse 通过追溯码查单据 API返回值
// alibaba.alihealth.drug.kyt.codetobill
//
// 通过追溯码查单据
type AlibabaAlihealthDrugKytCodetobillAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytCodetobillAPIResponseModel
}

// AlibabaAlihealthDrugKytCodetobillAPIResponseModel is 通过追溯码查单据 成功返回结果
type AlibabaAlihealthDrugKytCodetobillAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_codetobill_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytCodetobillResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
