package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardReserveAPIResponse 工单预约 API返回值
// tmall.servicecenter.workcard.reserve
//
// 服务工单更新通用接口
type TmallServicecenterWorkcardReserveAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardReserveAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardReserveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterWorkcardReserveAPIResponseModel).Reset()
}

// TmallServicecenterWorkcardReserveAPIResponseModel is 工单预约 成功返回结果
type TmallServicecenterWorkcardReserveAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_reserve_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *TmallServicecenterWorkcardReserveResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardReserveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterWorkcardReserveAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardReserveAPIResponse)
	},
}

// GetTmallServicecenterWorkcardReserveAPIResponse 从 sync.Pool 获取 TmallServicecenterWorkcardReserveAPIResponse
func GetTmallServicecenterWorkcardReserveAPIResponse() *TmallServicecenterWorkcardReserveAPIResponse {
	return poolTmallServicecenterWorkcardReserveAPIResponse.Get().(*TmallServicecenterWorkcardReserveAPIResponse)
}

// ReleaseTmallServicecenterWorkcardReserveAPIResponse 将 TmallServicecenterWorkcardReserveAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterWorkcardReserveAPIResponse(v *TmallServicecenterWorkcardReserveAPIResponse) {
	v.Reset()
	poolTmallServicecenterWorkcardReserveAPIResponse.Put(v)
}
