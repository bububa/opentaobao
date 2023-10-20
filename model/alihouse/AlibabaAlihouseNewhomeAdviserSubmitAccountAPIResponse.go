package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeAdviserSubmitAccountAPIResponse 顾问入职离职 API返回值
// alibaba.alihouse.newhome.adviser.submit.account
//
// 顾问入职离职
type AlibabaAlihouseNewhomeAdviserSubmitAccountAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeAdviserSubmitAccountAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeAdviserSubmitAccountAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeAdviserSubmitAccountAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeAdviserSubmitAccountAPIResponseModel is 顾问入职离职 成功返回结果
type AlibabaAlihouseNewhomeAdviserSubmitAccountAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_adviser_submit_account_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeAdviserSubmitAccountResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeAdviserSubmitAccountAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeAdviserSubmitAccountAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeAdviserSubmitAccountAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeAdviserSubmitAccountAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeAdviserSubmitAccountAPIResponse
func GetAlibabaAlihouseNewhomeAdviserSubmitAccountAPIResponse() *AlibabaAlihouseNewhomeAdviserSubmitAccountAPIResponse {
	return poolAlibabaAlihouseNewhomeAdviserSubmitAccountAPIResponse.Get().(*AlibabaAlihouseNewhomeAdviserSubmitAccountAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeAdviserSubmitAccountAPIResponse 将 AlibabaAlihouseNewhomeAdviserSubmitAccountAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeAdviserSubmitAccountAPIResponse(v *AlibabaAlihouseNewhomeAdviserSubmitAccountAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeAdviserSubmitAccountAPIResponse.Put(v)
}
