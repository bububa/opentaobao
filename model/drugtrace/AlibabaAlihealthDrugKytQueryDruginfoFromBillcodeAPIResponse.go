package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIResponse 根据单据编号查询单据明细 API返回值
// alibaba.alihealth.drug.kyt.query.druginfo.from.billcode
//
// 根据单据编号查询单据明细
type AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIResponseModel is 根据单据编号查询单据明细 成功返回结果
type AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_query_druginfo_from_billcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIResponse
func GetAlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIResponse() *AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIResponse {
	return poolAlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIResponse.Get().(*AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIResponse 将 AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIResponse(v *AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIResponse.Put(v)
}
