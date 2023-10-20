package tmallsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterAnomalyrecourseHomedecorationResponseAPIResponse 天猫服务平台商家投诉单服务商响应接口 API返回值
// tmall.servicecenter.anomalyrecourse.homedecoration.response
//
// 天猫服务平台商家投诉单服务商响应接口
type TmallServicecenterAnomalyrecourseHomedecorationResponseAPIResponse struct {
	model.CommonResponse
	TmallServicecenterAnomalyrecourseHomedecorationResponseAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterAnomalyrecourseHomedecorationResponseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterAnomalyrecourseHomedecorationResponseAPIResponseModel).Reset()
}

// TmallServicecenterAnomalyrecourseHomedecorationResponseAPIResponseModel is 天猫服务平台商家投诉单服务商响应接口 成功返回结果
type TmallServicecenterAnomalyrecourseHomedecorationResponseAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_anomalyrecourse_homedecoration_response_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallServicecenterAnomalyrecourseHomedecorationResponseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterAnomalyrecourseHomedecorationResponseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterAnomalyrecourseHomedecorationResponseAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterAnomalyrecourseHomedecorationResponseAPIResponse)
	},
}

// GetTmallServicecenterAnomalyrecourseHomedecorationResponseAPIResponse 从 sync.Pool 获取 TmallServicecenterAnomalyrecourseHomedecorationResponseAPIResponse
func GetTmallServicecenterAnomalyrecourseHomedecorationResponseAPIResponse() *TmallServicecenterAnomalyrecourseHomedecorationResponseAPIResponse {
	return poolTmallServicecenterAnomalyrecourseHomedecorationResponseAPIResponse.Get().(*TmallServicecenterAnomalyrecourseHomedecorationResponseAPIResponse)
}

// ReleaseTmallServicecenterAnomalyrecourseHomedecorationResponseAPIResponse 将 TmallServicecenterAnomalyrecourseHomedecorationResponseAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterAnomalyrecourseHomedecorationResponseAPIResponse(v *TmallServicecenterAnomalyrecourseHomedecorationResponseAPIResponse) {
	v.Reset()
	poolTmallServicecenterAnomalyrecourseHomedecorationResponseAPIResponse.Put(v)
}
