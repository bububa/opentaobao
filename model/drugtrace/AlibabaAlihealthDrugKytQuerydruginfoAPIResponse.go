package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytQuerydruginfoAPIResponse
码查询药品 API返回值
alibaba.alihealth.drug.kyt.querydruginfo

通过追溯码查询药品信息 */
type AlibabaAlihealthDrugKytQuerydruginfoAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytQuerydruginfoAPIResponseModel
}

// AlibabaAlihealthDrugKytQuerydruginfoAPIResponseModel is 码查询药品 成功返回结果
type AlibabaAlihealthDrugKytQuerydruginfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_querydruginfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaAlihealthDrugKytQuerydruginfoResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
