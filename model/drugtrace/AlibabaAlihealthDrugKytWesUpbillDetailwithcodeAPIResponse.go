package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytWesUpbillDetailwithcodeAPIResponse 查询上游出库单明细（带追溯码信息） API返回值
// alibaba.alihealth.drug.kyt.wes.upbill.detailwithcode
//
// 查询上游出库单明细(带追溯码信息)
type AlibabaAlihealthDrugKytWesUpbillDetailwithcodeAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytWesUpbillDetailwithcodeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytWesUpbillDetailwithcodeAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytWesUpbillDetailwithcodeAPIResponseModel is 查询上游出库单明细（带追溯码信息） 成功返回结果
type AlibabaAlihealthDrugKytWesUpbillDetailwithcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_wes_upbill_detailwithcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytWesUpbillDetailwithcodeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytWesUpbillDetailwithcodeAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytWesUpbillDetailwithcodeAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytWesUpbillDetailwithcodeAPIResponse
func GetAlibabaAlihealthDrugKytWesUpbillDetailwithcodeAPIResponse() *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeAPIResponse {
	return poolAlibabaAlihealthDrugKytWesUpbillDetailwithcodeAPIResponse.Get().(*AlibabaAlihealthDrugKytWesUpbillDetailwithcodeAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytWesUpbillDetailwithcodeAPIResponse 将 AlibabaAlihealthDrugKytWesUpbillDetailwithcodeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytWesUpbillDetailwithcodeAPIResponse(v *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytWesUpbillDetailwithcodeAPIResponse.Put(v)
}
