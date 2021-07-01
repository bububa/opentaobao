package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugCodeCommonListCodeinfoAPIResponse
通用查询码接口 API返回值
alibaba.alihealth.drug.code.common.list.codeinfo

通用查询码接口 */
type AlibabaAlihealthDrugCodeCommonListCodeinfoAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugCodeCommonListCodeinfoAPIResponseModel
}

// AlibabaAlihealthDrugCodeCommonListCodeinfoAPIResponseModel is 通用查询码接口 成功返回结果
type AlibabaAlihealthDrugCodeCommonListCodeinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_code_common_list_codeinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 最外层结果
	Result *AlibabaAlihealthDrugCodeCommonListCodeinfoResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
