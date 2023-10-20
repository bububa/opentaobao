package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytWesDrugrescodeAPIResponse 查询药品码段信息 API返回值
// alibaba.alihealth.drug.kyt.wes.drugrescode
//
// 查询药品码段信息
type AlibabaAlihealthDrugKytWesDrugrescodeAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytWesDrugrescodeAPIResponseModel
}

// AlibabaAlihealthDrugKytWesDrugrescodeAPIResponseModel is 查询药品码段信息 成功返回结果
type AlibabaAlihealthDrugKytWesDrugrescodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_wes_drugrescode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytWesDrugrescodeResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
