package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytWesSearchbillDetailAPIResponse 查询单据详情 API返回值
// alibaba.alihealth.drug.kyt.wes.searchbill.detail
//
// 根据单据号码查询码单据详情和码信息
type AlibabaAlihealthDrugKytWesSearchbillDetailAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytWesSearchbillDetailAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytWesSearchbillDetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytWesSearchbillDetailAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytWesSearchbillDetailAPIResponseModel is 查询单据详情 成功返回结果
type AlibabaAlihealthDrugKytWesSearchbillDetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_wes_searchbill_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytWesSearchbillDetailResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytWesSearchbillDetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytWesSearchbillDetailAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytWesSearchbillDetailAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytWesSearchbillDetailAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytWesSearchbillDetailAPIResponse
func GetAlibabaAlihealthDrugKytWesSearchbillDetailAPIResponse() *AlibabaAlihealthDrugKytWesSearchbillDetailAPIResponse {
	return poolAlibabaAlihealthDrugKytWesSearchbillDetailAPIResponse.Get().(*AlibabaAlihealthDrugKytWesSearchbillDetailAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytWesSearchbillDetailAPIResponse 将 AlibabaAlihealthDrugKytWesSearchbillDetailAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytWesSearchbillDetailAPIResponse(v *AlibabaAlihealthDrugKytWesSearchbillDetailAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytWesSearchbillDetailAPIResponse.Put(v)
}
