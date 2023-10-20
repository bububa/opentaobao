package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIResponse 根据单据编号查询单据明细 API返回值
// alibaba.alihealth.drug.kyt.query.specia.vaccin.billcode
//
// 根据单据编号查询单据明细
type AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIResponseModel is 根据单据编号查询单据明细 成功返回结果
type AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_query_specia_vaccin_billcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIResponse
func GetAlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIResponse() *AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIResponse {
	return poolAlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIResponse.Get().(*AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIResponse 将 AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIResponse(v *AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIResponse.Put(v)
}
