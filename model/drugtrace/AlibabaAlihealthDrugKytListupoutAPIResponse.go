package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytListupoutAPIResponse 查询货主/本企业上游企业出库单据信息 API返回值
// alibaba.alihealth.drug.kyt.listupout
//
// 查询货主/本企业上游企业出库单据信息
type AlibabaAlihealthDrugKytListupoutAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytListupoutAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytListupoutAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytListupoutAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytListupoutAPIResponseModel is 查询货主/本企业上游企业出库单据信息 成功返回结果
type AlibabaAlihealthDrugKytListupoutAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_listupout_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytListupoutResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytListupoutAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytListupoutAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytListupoutAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytListupoutAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytListupoutAPIResponse
func GetAlibabaAlihealthDrugKytListupoutAPIResponse() *AlibabaAlihealthDrugKytListupoutAPIResponse {
	return poolAlibabaAlihealthDrugKytListupoutAPIResponse.Get().(*AlibabaAlihealthDrugKytListupoutAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytListupoutAPIResponse 将 AlibabaAlihealthDrugKytListupoutAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytListupoutAPIResponse(v *AlibabaAlihealthDrugKytListupoutAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytListupoutAPIResponse.Put(v)
}
