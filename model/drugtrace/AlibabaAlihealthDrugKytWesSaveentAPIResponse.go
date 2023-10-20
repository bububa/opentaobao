package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytWesSaveentAPIResponse 新增往来单位企业记录 API返回值
// alibaba.alihealth.drug.kyt.wes.saveent
//
// 新增往来单位企业记录
type AlibabaAlihealthDrugKytWesSaveentAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytWesSaveentAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytWesSaveentAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytWesSaveentAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytWesSaveentAPIResponseModel is 新增往来单位企业记录 成功返回结果
type AlibabaAlihealthDrugKytWesSaveentAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_wes_saveent_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 往来单位新增接口返回
	Result *AlibabaAlihealthDrugKytWesSaveentResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytWesSaveentAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytWesSaveentAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytWesSaveentAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytWesSaveentAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytWesSaveentAPIResponse
func GetAlibabaAlihealthDrugKytWesSaveentAPIResponse() *AlibabaAlihealthDrugKytWesSaveentAPIResponse {
	return poolAlibabaAlihealthDrugKytWesSaveentAPIResponse.Get().(*AlibabaAlihealthDrugKytWesSaveentAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytWesSaveentAPIResponse 将 AlibabaAlihealthDrugKytWesSaveentAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytWesSaveentAPIResponse(v *AlibabaAlihealthDrugKytWesSaveentAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytWesSaveentAPIResponse.Put(v)
}
