package tmallsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterAnomalyrecourseRemarkUpdateAPIResponse 天猫服务平台一键求助单服务商备注更新接口 API返回值
// tmall.servicecenter.anomalyrecourse.remark.update
//
// 一键求助服务商可以回传备注
type TmallServicecenterAnomalyrecourseRemarkUpdateAPIResponse struct {
	model.CommonResponse
	TmallServicecenterAnomalyrecourseRemarkUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterAnomalyrecourseRemarkUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterAnomalyrecourseRemarkUpdateAPIResponseModel).Reset()
}

// TmallServicecenterAnomalyrecourseRemarkUpdateAPIResponseModel is 天猫服务平台一键求助单服务商备注更新接口 成功返回结果
type TmallServicecenterAnomalyrecourseRemarkUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_anomalyrecourse_remark_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterAnomalyrecourseRemarkUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTmallServicecenterAnomalyrecourseRemarkUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterAnomalyrecourseRemarkUpdateAPIResponse)
	},
}

// GetTmallServicecenterAnomalyrecourseRemarkUpdateAPIResponse 从 sync.Pool 获取 TmallServicecenterAnomalyrecourseRemarkUpdateAPIResponse
func GetTmallServicecenterAnomalyrecourseRemarkUpdateAPIResponse() *TmallServicecenterAnomalyrecourseRemarkUpdateAPIResponse {
	return poolTmallServicecenterAnomalyrecourseRemarkUpdateAPIResponse.Get().(*TmallServicecenterAnomalyrecourseRemarkUpdateAPIResponse)
}

// ReleaseTmallServicecenterAnomalyrecourseRemarkUpdateAPIResponse 将 TmallServicecenterAnomalyrecourseRemarkUpdateAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterAnomalyrecourseRemarkUpdateAPIResponse(v *TmallServicecenterAnomalyrecourseRemarkUpdateAPIResponse) {
	v.Reset()
	poolTmallServicecenterAnomalyrecourseRemarkUpdateAPIResponse.Put(v)
}
