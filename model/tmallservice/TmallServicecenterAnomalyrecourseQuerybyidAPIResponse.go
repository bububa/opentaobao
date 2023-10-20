package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterAnomalyrecourseQuerybyidAPIResponse 根据一键求助id查询指定服务商的一键求助单 API返回值
// tmall.servicecenter.anomalyrecourse.querybyid
//
// 根据一键求助id查询指定服务商的一键求助单
type TmallServicecenterAnomalyrecourseQuerybyidAPIResponse struct {
	model.CommonResponse
	TmallServicecenterAnomalyrecourseQuerybyidAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterAnomalyrecourseQuerybyidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterAnomalyrecourseQuerybyidAPIResponseModel).Reset()
}

// TmallServicecenterAnomalyrecourseQuerybyidAPIResponseModel is 根据一键求助id查询指定服务商的一键求助单 成功返回结果
type TmallServicecenterAnomalyrecourseQuerybyidAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_anomalyrecourse_querybyid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterAnomalyrecourseQuerybyidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterAnomalyrecourseQuerybyidAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterAnomalyrecourseQuerybyidAPIResponse)
	},
}

// GetTmallServicecenterAnomalyrecourseQuerybyidAPIResponse 从 sync.Pool 获取 TmallServicecenterAnomalyrecourseQuerybyidAPIResponse
func GetTmallServicecenterAnomalyrecourseQuerybyidAPIResponse() *TmallServicecenterAnomalyrecourseQuerybyidAPIResponse {
	return poolTmallServicecenterAnomalyrecourseQuerybyidAPIResponse.Get().(*TmallServicecenterAnomalyrecourseQuerybyidAPIResponse)
}

// ReleaseTmallServicecenterAnomalyrecourseQuerybyidAPIResponse 将 TmallServicecenterAnomalyrecourseQuerybyidAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterAnomalyrecourseQuerybyidAPIResponse(v *TmallServicecenterAnomalyrecourseQuerybyidAPIResponse) {
	v.Reset()
	poolTmallServicecenterAnomalyrecourseQuerybyidAPIResponse.Put(v)
}
