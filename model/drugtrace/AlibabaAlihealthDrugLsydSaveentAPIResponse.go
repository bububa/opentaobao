package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugLsydSaveentAPIResponse 零售药店往来单位新增 API返回值
// alibaba.alihealth.drug.lsyd.saveent
//
// 新增往来单位企业记录
type AlibabaAlihealthDrugLsydSaveentAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugLsydSaveentAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugLsydSaveentAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugLsydSaveentAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugLsydSaveentAPIResponseModel is 零售药店往来单位新增 成功返回结果
type AlibabaAlihealthDrugLsydSaveentAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_lsyd_saveent_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 往来单位新增接口返回
	Result *AlibabaAlihealthDrugLsydSaveentResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugLsydSaveentAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugLsydSaveentAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugLsydSaveentAPIResponse)
	},
}

// GetAlibabaAlihealthDrugLsydSaveentAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugLsydSaveentAPIResponse
func GetAlibabaAlihealthDrugLsydSaveentAPIResponse() *AlibabaAlihealthDrugLsydSaveentAPIResponse {
	return poolAlibabaAlihealthDrugLsydSaveentAPIResponse.Get().(*AlibabaAlihealthDrugLsydSaveentAPIResponse)
}

// ReleaseAlibabaAlihealthDrugLsydSaveentAPIResponse 将 AlibabaAlihealthDrugLsydSaveentAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugLsydSaveentAPIResponse(v *AlibabaAlihealthDrugLsydSaveentAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugLsydSaveentAPIResponse.Put(v)
}
