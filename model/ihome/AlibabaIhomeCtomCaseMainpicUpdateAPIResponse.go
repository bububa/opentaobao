package ihome

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIhomeCtomCaseMainpicUpdateAPIResponse 方案渲染图修改 API返回值
// alibaba.ihome.ctom.case.mainpic.update
//
// 用于在门店工作台里的编辑器保存方案，由三维家后端调用阿里后端，保存方案信息
// 此接口只允许ihome业务使用，用于门店的编辑功能，只允许广东三维家信息科技有限公司一家公司调用，不适用于其他业务。
type AlibabaIhomeCtomCaseMainpicUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaIhomeCtomCaseMainpicUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIhomeCtomCaseMainpicUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIhomeCtomCaseMainpicUpdateAPIResponseModel).Reset()
}

// AlibabaIhomeCtomCaseMainpicUpdateAPIResponseModel is 方案渲染图修改 成功返回结果
type AlibabaIhomeCtomCaseMainpicUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ihome_ctom_case_mainpic_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 根据站点名称查询产品
	ApiResult *AlibabaIhomeCtomCaseMainpicUpdateApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIhomeCtomCaseMainpicUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaIhomeCtomCaseMainpicUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIhomeCtomCaseMainpicUpdateAPIResponse)
	},
}

// GetAlibabaIhomeCtomCaseMainpicUpdateAPIResponse 从 sync.Pool 获取 AlibabaIhomeCtomCaseMainpicUpdateAPIResponse
func GetAlibabaIhomeCtomCaseMainpicUpdateAPIResponse() *AlibabaIhomeCtomCaseMainpicUpdateAPIResponse {
	return poolAlibabaIhomeCtomCaseMainpicUpdateAPIResponse.Get().(*AlibabaIhomeCtomCaseMainpicUpdateAPIResponse)
}

// ReleaseAlibabaIhomeCtomCaseMainpicUpdateAPIResponse 将 AlibabaIhomeCtomCaseMainpicUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIhomeCtomCaseMainpicUpdateAPIResponse(v *AlibabaIhomeCtomCaseMainpicUpdateAPIResponse) {
	v.Reset()
	poolAlibabaIhomeCtomCaseMainpicUpdateAPIResponse.Put(v)
}
