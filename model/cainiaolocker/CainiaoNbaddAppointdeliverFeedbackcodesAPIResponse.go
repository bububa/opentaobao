package cainiaolocker

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoNbaddAppointdeliverFeedbackcodesAPIResponse 服务质量反馈编码列表 API返回值
// cainiao.nbadd.appointdeliver.feedbackcodes
//
// 服务质量反馈编码列表，建议获取数据后缓存在本地，定时刷新即可
type CainiaoNbaddAppointdeliverFeedbackcodesAPIResponse struct {
	model.CommonResponse
	CainiaoNbaddAppointdeliverFeedbackcodesAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoNbaddAppointdeliverFeedbackcodesAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoNbaddAppointdeliverFeedbackcodesAPIResponseModel).Reset()
}

// CainiaoNbaddAppointdeliverFeedbackcodesAPIResponseModel is 服务质量反馈编码列表 成功返回结果
type CainiaoNbaddAppointdeliverFeedbackcodesAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_nbadd_appointdeliver_feedbackcodes_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的具体数据
	ResultList []FeedbackCodeDto `json:"result_list,omitempty" xml:"result_list>feedback_code_dto,omitempty"`
	// 错误描述
	ResultDesc string `json:"result_desc,omitempty" xml:"result_desc,omitempty"`
	// 错误编码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 接口调用是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoNbaddAppointdeliverFeedbackcodesAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultList = m.ResultList[:0]
	m.ResultDesc = ""
	m.ResultCode = ""
	m.IsSuccess = false
}

var poolCainiaoNbaddAppointdeliverFeedbackcodesAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoNbaddAppointdeliverFeedbackcodesAPIResponse)
	},
}

// GetCainiaoNbaddAppointdeliverFeedbackcodesAPIResponse 从 sync.Pool 获取 CainiaoNbaddAppointdeliverFeedbackcodesAPIResponse
func GetCainiaoNbaddAppointdeliverFeedbackcodesAPIResponse() *CainiaoNbaddAppointdeliverFeedbackcodesAPIResponse {
	return poolCainiaoNbaddAppointdeliverFeedbackcodesAPIResponse.Get().(*CainiaoNbaddAppointdeliverFeedbackcodesAPIResponse)
}

// ReleaseCainiaoNbaddAppointdeliverFeedbackcodesAPIResponse 将 CainiaoNbaddAppointdeliverFeedbackcodesAPIResponse 保存到 sync.Pool
func ReleaseCainiaoNbaddAppointdeliverFeedbackcodesAPIResponse(v *CainiaoNbaddAppointdeliverFeedbackcodesAPIResponse) {
	v.Reset()
	poolCainiaoNbaddAppointdeliverFeedbackcodesAPIResponse.Put(v)
}
