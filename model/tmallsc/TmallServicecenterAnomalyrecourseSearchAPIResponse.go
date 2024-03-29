package tmallsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterAnomalyrecourseSearchAPIResponse 天猫服务平台服务商一键求助单查询 API返回值
// tmall.servicecenter.anomalyrecourse.search
//
// 天猫服务平台服务商一键求助单查询
type TmallServicecenterAnomalyrecourseSearchAPIResponse struct {
	model.CommonResponse
	TmallServicecenterAnomalyrecourseSearchAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterAnomalyrecourseSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterAnomalyrecourseSearchAPIResponseModel).Reset()
}

// TmallServicecenterAnomalyrecourseSearchAPIResponseModel is 天猫服务平台服务商一键求助单查询 成功返回结果
type TmallServicecenterAnomalyrecourseSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_anomalyrecourse_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterAnomalyrecourseSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterAnomalyrecourseSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterAnomalyrecourseSearchAPIResponse)
	},
}

// GetTmallServicecenterAnomalyrecourseSearchAPIResponse 从 sync.Pool 获取 TmallServicecenterAnomalyrecourseSearchAPIResponse
func GetTmallServicecenterAnomalyrecourseSearchAPIResponse() *TmallServicecenterAnomalyrecourseSearchAPIResponse {
	return poolTmallServicecenterAnomalyrecourseSearchAPIResponse.Get().(*TmallServicecenterAnomalyrecourseSearchAPIResponse)
}

// ReleaseTmallServicecenterAnomalyrecourseSearchAPIResponse 将 TmallServicecenterAnomalyrecourseSearchAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterAnomalyrecourseSearchAPIResponse(v *TmallServicecenterAnomalyrecourseSearchAPIResponse) {
	v.Reset()
	poolTmallServicecenterAnomalyrecourseSearchAPIResponse.Put(v)
}
