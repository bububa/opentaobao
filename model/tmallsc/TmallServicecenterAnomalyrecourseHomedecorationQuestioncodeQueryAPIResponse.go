package tmallsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIResponse 天猫服务平台商家投诉单问题列表查询 API返回值
// tmall.servicecenter.anomalyrecourse.homedecoration.questioncode.query
//
// 天猫服务平台商家投诉单问题列表查询
type TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIResponse struct {
	model.CommonResponse
	TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIResponseModel).Reset()
}

// TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIResponseModel is 天猫服务平台商家投诉单问题列表查询 成功返回结果
type TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_anomalyrecourse_homedecoration_questioncode_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIResponse)
	},
}

// GetTmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIResponse 从 sync.Pool 获取 TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIResponse
func GetTmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIResponse() *TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIResponse {
	return poolTmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIResponse.Get().(*TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIResponse)
}

// ReleaseTmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIResponse 将 TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIResponse(v *TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIResponse) {
	v.Reset()
	poolTmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIResponse.Put(v)
}
