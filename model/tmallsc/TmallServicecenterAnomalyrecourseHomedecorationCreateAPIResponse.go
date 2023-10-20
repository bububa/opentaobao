package tmallsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterAnomalyrecourseHomedecorationCreateAPIResponse 天猫服务平台服务商代商家发起投诉单 API返回值
// tmall.servicecenter.anomalyrecourse.homedecoration.create
//
// 天猫服务平台服务商代商家发起投诉单
type TmallServicecenterAnomalyrecourseHomedecorationCreateAPIResponse struct {
	model.CommonResponse
	TmallServicecenterAnomalyrecourseHomedecorationCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterAnomalyrecourseHomedecorationCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterAnomalyrecourseHomedecorationCreateAPIResponseModel).Reset()
}

// TmallServicecenterAnomalyrecourseHomedecorationCreateAPIResponseModel is 天猫服务平台服务商代商家发起投诉单 成功返回结果
type TmallServicecenterAnomalyrecourseHomedecorationCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_anomalyrecourse_homedecoration_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallServicecenterAnomalyrecourseHomedecorationCreateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterAnomalyrecourseHomedecorationCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterAnomalyrecourseHomedecorationCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterAnomalyrecourseHomedecorationCreateAPIResponse)
	},
}

// GetTmallServicecenterAnomalyrecourseHomedecorationCreateAPIResponse 从 sync.Pool 获取 TmallServicecenterAnomalyrecourseHomedecorationCreateAPIResponse
func GetTmallServicecenterAnomalyrecourseHomedecorationCreateAPIResponse() *TmallServicecenterAnomalyrecourseHomedecorationCreateAPIResponse {
	return poolTmallServicecenterAnomalyrecourseHomedecorationCreateAPIResponse.Get().(*TmallServicecenterAnomalyrecourseHomedecorationCreateAPIResponse)
}

// ReleaseTmallServicecenterAnomalyrecourseHomedecorationCreateAPIResponse 将 TmallServicecenterAnomalyrecourseHomedecorationCreateAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterAnomalyrecourseHomedecorationCreateAPIResponse(v *TmallServicecenterAnomalyrecourseHomedecorationCreateAPIResponse) {
	v.Reset()
	poolTmallServicecenterAnomalyrecourseHomedecorationCreateAPIResponse.Put(v)
}
