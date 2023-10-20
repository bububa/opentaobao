package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytSaveentAPIResponse 新增往来单位企业 API返回值
// alibaba.alihealth.drug.kyt.saveent
//
// 新增往来单位企业记录
type AlibabaAlihealthDrugKytSaveentAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytSaveentAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytSaveentAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytSaveentAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytSaveentAPIResponseModel is 新增往来单位企业 成功返回结果
type AlibabaAlihealthDrugKytSaveentAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_saveent_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 往来单位新增接口返回
	Result *AlibabaAlihealthDrugKytSaveentResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytSaveentAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytSaveentAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytSaveentAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytSaveentAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytSaveentAPIResponse
func GetAlibabaAlihealthDrugKytSaveentAPIResponse() *AlibabaAlihealthDrugKytSaveentAPIResponse {
	return poolAlibabaAlihealthDrugKytSaveentAPIResponse.Get().(*AlibabaAlihealthDrugKytSaveentAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytSaveentAPIResponse 将 AlibabaAlihealthDrugKytSaveentAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytSaveentAPIResponse(v *AlibabaAlihealthDrugKytSaveentAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytSaveentAPIResponse.Put(v)
}
