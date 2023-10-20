package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytSearchbillAPIResponse 通过时间段批量查询入出库单信息 API返回值
// alibaba.alihealth.drug.kyt.searchbill
//
// 通过时间段批量查询入出库单信息
type AlibabaAlihealthDrugKytSearchbillAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytSearchbillAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytSearchbillAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytSearchbillAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytSearchbillAPIResponseModel is 通过时间段批量查询入出库单信息 成功返回结果
type AlibabaAlihealthDrugKytSearchbillAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_searchbill_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytSearchbillResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytSearchbillAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytSearchbillAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytSearchbillAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytSearchbillAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytSearchbillAPIResponse
func GetAlibabaAlihealthDrugKytSearchbillAPIResponse() *AlibabaAlihealthDrugKytSearchbillAPIResponse {
	return poolAlibabaAlihealthDrugKytSearchbillAPIResponse.Get().(*AlibabaAlihealthDrugKytSearchbillAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytSearchbillAPIResponse 将 AlibabaAlihealthDrugKytSearchbillAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytSearchbillAPIResponse(v *AlibabaAlihealthDrugKytSearchbillAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytSearchbillAPIResponse.Put(v)
}
