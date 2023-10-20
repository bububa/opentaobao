package tmallsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterAnomalyrecourseHomedecorationCloseAPIResponse 天猫服务平台商家投诉单服务商完结接口 API返回值
// tmall.servicecenter.anomalyrecourse.homedecoration.close
//
// 天猫服务平台商家投诉单服务商完结接口
type TmallServicecenterAnomalyrecourseHomedecorationCloseAPIResponse struct {
	model.CommonResponse
	TmallServicecenterAnomalyrecourseHomedecorationCloseAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterAnomalyrecourseHomedecorationCloseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterAnomalyrecourseHomedecorationCloseAPIResponseModel).Reset()
}

// TmallServicecenterAnomalyrecourseHomedecorationCloseAPIResponseModel is 天猫服务平台商家投诉单服务商完结接口 成功返回结果
type TmallServicecenterAnomalyrecourseHomedecorationCloseAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_anomalyrecourse_homedecoration_close_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallServicecenterAnomalyrecourseHomedecorationCloseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterAnomalyrecourseHomedecorationCloseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterAnomalyrecourseHomedecorationCloseAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterAnomalyrecourseHomedecorationCloseAPIResponse)
	},
}

// GetTmallServicecenterAnomalyrecourseHomedecorationCloseAPIResponse 从 sync.Pool 获取 TmallServicecenterAnomalyrecourseHomedecorationCloseAPIResponse
func GetTmallServicecenterAnomalyrecourseHomedecorationCloseAPIResponse() *TmallServicecenterAnomalyrecourseHomedecorationCloseAPIResponse {
	return poolTmallServicecenterAnomalyrecourseHomedecorationCloseAPIResponse.Get().(*TmallServicecenterAnomalyrecourseHomedecorationCloseAPIResponse)
}

// ReleaseTmallServicecenterAnomalyrecourseHomedecorationCloseAPIResponse 将 TmallServicecenterAnomalyrecourseHomedecorationCloseAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterAnomalyrecourseHomedecorationCloseAPIResponse(v *TmallServicecenterAnomalyrecourseHomedecorationCloseAPIResponse) {
	v.Reset()
	poolTmallServicecenterAnomalyrecourseHomedecorationCloseAPIResponse.Put(v)
}
