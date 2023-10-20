package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardReservefailAPIResponse 预约失败 API返回值
// tmall.servicecenter.workcard.reservefail
//
// 服务商调用该接口回传工单预约失败
type TmallServicecenterWorkcardReservefailAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardReservefailAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardReservefailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterWorkcardReservefailAPIResponseModel).Reset()
}

// TmallServicecenterWorkcardReservefailAPIResponseModel is 预约失败 成功返回结果
type TmallServicecenterWorkcardReservefailAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_reservefail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// -
	Result *TmallServicecenterWorkcardReservefailResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardReservefailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterWorkcardReservefailAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardReservefailAPIResponse)
	},
}

// GetTmallServicecenterWorkcardReservefailAPIResponse 从 sync.Pool 获取 TmallServicecenterWorkcardReservefailAPIResponse
func GetTmallServicecenterWorkcardReservefailAPIResponse() *TmallServicecenterWorkcardReservefailAPIResponse {
	return poolTmallServicecenterWorkcardReservefailAPIResponse.Get().(*TmallServicecenterWorkcardReservefailAPIResponse)
}

// ReleaseTmallServicecenterWorkcardReservefailAPIResponse 将 TmallServicecenterWorkcardReservefailAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterWorkcardReservefailAPIResponse(v *TmallServicecenterWorkcardReservefailAPIResponse) {
	v.Reset()
	poolTmallServicecenterWorkcardReservefailAPIResponse.Put(v)
}
