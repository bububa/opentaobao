package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIResponse
根据追溯码修改69码 API返回值
alibaba.alihealth.drug.updatebarcode.bytraccode

根据追溯码修改69码 */
type AlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIResponseModel
}

// AlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIResponseModel is 根据追溯码修改69码 成功返回结果
type AlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_updatebarcode_bytraccode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
}
