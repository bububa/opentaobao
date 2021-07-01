package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIResponse
扫码营销码查询 API返回值
alibaba.alihealth.drug.code.kyt.smyx.querycode

此接口针对有码药品，提供可通过追溯码获取该药品的基础信息和生产信息；
核查平台优先过滤非8开头的，长度非20位数字的码信息。 */
type AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIResponseModel
}

// AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIResponseModel is 扫码营销码查询 成功返回结果
type AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_code_kyt_smyx_querycode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 最外层结果
	Result *AlibabaAlihealthDrugCodeKytSmyxQuerycodeResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
