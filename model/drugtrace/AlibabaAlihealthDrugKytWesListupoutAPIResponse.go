package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytWesListupoutAPIResponse 查询货主/本企业上游企业出库单据信息 API返回值
// alibaba.alihealth.drug.kyt.wes.listupout
//
// 查询货主/本企业上游企业出库单据信息
type AlibabaAlihealthDrugKytWesListupoutAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytWesListupoutAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytWesListupoutAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytWesListupoutAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytWesListupoutAPIResponseModel is 查询货主/本企业上游企业出库单据信息 成功返回结果
type AlibabaAlihealthDrugKytWesListupoutAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_wes_listupout_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytWesListupoutResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytWesListupoutAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytWesListupoutAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytWesListupoutAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytWesListupoutAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytWesListupoutAPIResponse
func GetAlibabaAlihealthDrugKytWesListupoutAPIResponse() *AlibabaAlihealthDrugKytWesListupoutAPIResponse {
	return poolAlibabaAlihealthDrugKytWesListupoutAPIResponse.Get().(*AlibabaAlihealthDrugKytWesListupoutAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytWesListupoutAPIResponse 将 AlibabaAlihealthDrugKytWesListupoutAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytWesListupoutAPIResponse(v *AlibabaAlihealthDrugKytWesListupoutAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytWesListupoutAPIResponse.Put(v)
}
