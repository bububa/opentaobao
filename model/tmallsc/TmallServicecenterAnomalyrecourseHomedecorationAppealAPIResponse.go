package tmallsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterAnomalyrecourseHomedecorationAppealAPIResponse 天猫服务平台商家投诉单服务商申诉接口 API返回值
// tmall.servicecenter.anomalyrecourse.homedecoration.appeal
//
// 天猫服务平台商家投诉单服务商申诉接口
type TmallServicecenterAnomalyrecourseHomedecorationAppealAPIResponse struct {
	model.CommonResponse
	TmallServicecenterAnomalyrecourseHomedecorationAppealAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterAnomalyrecourseHomedecorationAppealAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterAnomalyrecourseHomedecorationAppealAPIResponseModel).Reset()
}

// TmallServicecenterAnomalyrecourseHomedecorationAppealAPIResponseModel is 天猫服务平台商家投诉单服务商申诉接口 成功返回结果
type TmallServicecenterAnomalyrecourseHomedecorationAppealAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_anomalyrecourse_homedecoration_appeal_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallServicecenterAnomalyrecourseHomedecorationAppealResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterAnomalyrecourseHomedecorationAppealAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterAnomalyrecourseHomedecorationAppealAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterAnomalyrecourseHomedecorationAppealAPIResponse)
	},
}

// GetTmallServicecenterAnomalyrecourseHomedecorationAppealAPIResponse 从 sync.Pool 获取 TmallServicecenterAnomalyrecourseHomedecorationAppealAPIResponse
func GetTmallServicecenterAnomalyrecourseHomedecorationAppealAPIResponse() *TmallServicecenterAnomalyrecourseHomedecorationAppealAPIResponse {
	return poolTmallServicecenterAnomalyrecourseHomedecorationAppealAPIResponse.Get().(*TmallServicecenterAnomalyrecourseHomedecorationAppealAPIResponse)
}

// ReleaseTmallServicecenterAnomalyrecourseHomedecorationAppealAPIResponse 将 TmallServicecenterAnomalyrecourseHomedecorationAppealAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterAnomalyrecourseHomedecorationAppealAPIResponse(v *TmallServicecenterAnomalyrecourseHomedecorationAppealAPIResponse) {
	v.Reset()
	poolTmallServicecenterAnomalyrecourseHomedecorationAppealAPIResponse.Put(v)
}
