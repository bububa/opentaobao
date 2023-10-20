package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytSearchbillDetailAPIResponse 查询单据详情 API返回值
// alibaba.alihealth.drug.kyt.searchbill.detail
//
// 根据单据号码查询码单据详情和码信息
type AlibabaAlihealthDrugKytSearchbillDetailAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytSearchbillDetailAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytSearchbillDetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytSearchbillDetailAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytSearchbillDetailAPIResponseModel is 查询单据详情 成功返回结果
type AlibabaAlihealthDrugKytSearchbillDetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_searchbill_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytSearchbillDetailResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytSearchbillDetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytSearchbillDetailAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytSearchbillDetailAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytSearchbillDetailAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytSearchbillDetailAPIResponse
func GetAlibabaAlihealthDrugKytSearchbillDetailAPIResponse() *AlibabaAlihealthDrugKytSearchbillDetailAPIResponse {
	return poolAlibabaAlihealthDrugKytSearchbillDetailAPIResponse.Get().(*AlibabaAlihealthDrugKytSearchbillDetailAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytSearchbillDetailAPIResponse 将 AlibabaAlihealthDrugKytSearchbillDetailAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytSearchbillDetailAPIResponse(v *AlibabaAlihealthDrugKytSearchbillDetailAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytSearchbillDetailAPIResponse.Put(v)
}
