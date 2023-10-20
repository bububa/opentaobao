package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeKytWesQuerycodeAPIResponse wes查询追溯码信息 API返回值
// alibaba.alihealth.drug.code.kyt.wes.querycode
//
// 此接口针对有码药品，提供可通过追溯码获取该药品的基础信息和生产信息；
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
type AlibabaAlihealthDrugCodeKytWesQuerycodeAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugCodeKytWesQuerycodeAPIResponseModel
}

// AlibabaAlihealthDrugCodeKytWesQuerycodeAPIResponseModel is wes查询追溯码信息 成功返回结果
type AlibabaAlihealthDrugCodeKytWesQuerycodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_code_kyt_wes_querycode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 最外层结果
	Result *AlibabaAlihealthDrugCodeKytWesQuerycodeResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
