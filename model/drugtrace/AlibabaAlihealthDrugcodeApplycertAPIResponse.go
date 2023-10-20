package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcodeApplycertAPIResponse 申请证书为对接方 API返回值
// alibaba.alihealth.drugcode.applycert
//
// 申请证书 为对接方(当前是药厂和中心化系统)
type AlibabaAlihealthDrugcodeApplycertAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugcodeApplycertAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugcodeApplycertAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugcodeApplycertAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugcodeApplycertAPIResponseModel is 申请证书为对接方 成功返回结果
type AlibabaAlihealthDrugcodeApplycertAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugcode_applycert_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugcodeApplycertAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugcodeApplycertAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugcodeApplycertAPIResponse)
	},
}

// GetAlibabaAlihealthDrugcodeApplycertAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugcodeApplycertAPIResponse
func GetAlibabaAlihealthDrugcodeApplycertAPIResponse() *AlibabaAlihealthDrugcodeApplycertAPIResponse {
	return poolAlibabaAlihealthDrugcodeApplycertAPIResponse.Get().(*AlibabaAlihealthDrugcodeApplycertAPIResponse)
}

// ReleaseAlibabaAlihealthDrugcodeApplycertAPIResponse 将 AlibabaAlihealthDrugcodeApplycertAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugcodeApplycertAPIResponse(v *AlibabaAlihealthDrugcodeApplycertAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugcodeApplycertAPIResponse.Put(v)
}
