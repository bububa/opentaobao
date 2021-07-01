package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytYyDrugcodesAPIResponse
查询药品是否赋码 API返回值
alibaba.alihealth.drug.kyt.yy.drugcodes

药品是否赋码 */
type AlibabaAlihealthDrugKytYyDrugcodesAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytYyDrugcodesAPIResponseModel
}

// AlibabaAlihealthDrugKytYyDrugcodesAPIResponseModel is 查询药品是否赋码 成功返回结果
type AlibabaAlihealthDrugKytYyDrugcodesAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_yy_drugcodes_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytYyDrugcodesResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
