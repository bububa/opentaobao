package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIResponse 天猫服务平台服务商查询商家投诉单 API返回值
// tmall.servicecenter.anomalyrecourse.homedecoration.querybyid
//
// 天猫服务平台服务商查询商家投诉单
type TmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIResponse struct {
	model.CommonResponse
	TmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIResponseModel).Reset()
}

// TmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIResponseModel is 天猫服务平台服务商查询商家投诉单 成功返回结果
type TmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_anomalyrecourse_homedecoration_querybyid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallServicecenterAnomalyrecourseHomedecorationQuerybyidResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIResponse)
	},
}

// GetTmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIResponse 从 sync.Pool 获取 TmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIResponse
func GetTmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIResponse() *TmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIResponse {
	return poolTmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIResponse.Get().(*TmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIResponse)
}

// ReleaseTmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIResponse 将 TmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIResponse(v *TmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIResponse) {
	v.Reset()
	poolTmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIResponse.Put(v)
}
