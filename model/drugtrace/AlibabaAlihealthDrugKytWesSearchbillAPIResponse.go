package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytWesSearchbillAPIResponse 通过时间段批量查询入出库单信息 API返回值
// alibaba.alihealth.drug.kyt.wes.searchbill
//
// 通过时间段批量查询入出库单信息
type AlibabaAlihealthDrugKytWesSearchbillAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytWesSearchbillAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytWesSearchbillAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytWesSearchbillAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytWesSearchbillAPIResponseModel is 通过时间段批量查询入出库单信息 成功返回结果
type AlibabaAlihealthDrugKytWesSearchbillAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_wes_searchbill_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytWesSearchbillResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytWesSearchbillAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytWesSearchbillAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytWesSearchbillAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytWesSearchbillAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytWesSearchbillAPIResponse
func GetAlibabaAlihealthDrugKytWesSearchbillAPIResponse() *AlibabaAlihealthDrugKytWesSearchbillAPIResponse {
	return poolAlibabaAlihealthDrugKytWesSearchbillAPIResponse.Get().(*AlibabaAlihealthDrugKytWesSearchbillAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytWesSearchbillAPIResponse 将 AlibabaAlihealthDrugKytWesSearchbillAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytWesSearchbillAPIResponse(v *AlibabaAlihealthDrugKytWesSearchbillAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytWesSearchbillAPIResponse.Put(v)
}
