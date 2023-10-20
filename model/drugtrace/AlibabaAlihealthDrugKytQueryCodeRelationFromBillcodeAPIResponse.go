package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIResponse 根据单据号码查询码单据详情和码信息 API返回值
// alibaba.alihealth.drug.kyt.query.code.relation.from.billcode
//
// 根据单据号码查询码单据详情和码信息
type AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIResponseModel is 根据单据号码查询码单据详情和码信息 成功返回结果
type AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_query_code_relation_from_billcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIResponse
func GetAlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIResponse() *AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIResponse {
	return poolAlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIResponse.Get().(*AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIResponse 将 AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIResponse(v *AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIResponse.Put(v)
}
