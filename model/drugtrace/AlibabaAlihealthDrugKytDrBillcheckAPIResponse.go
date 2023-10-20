package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDrBillcheckAPIResponse 疫苗追溯验证 API返回值
// alibaba.alihealth.drug.kyt.dr.billcheck
//
// 各级疾控在入库完成后，需要做追溯信息验证
type AlibabaAlihealthDrugKytDrBillcheckAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytDrBillcheckAPIResponseModel
}

// AlibabaAlihealthDrugKytDrBillcheckAPIResponseModel is 疫苗追溯验证 成功返回结果
type AlibabaAlihealthDrugKytDrBillcheckAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_dr_billcheck_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihealthDrugKytDrBillcheckResult `json:"result,omitempty" xml:"result,omitempty"`
}
