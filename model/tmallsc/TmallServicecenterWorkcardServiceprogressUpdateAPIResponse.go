package tmallsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardServiceprogressUpdateAPIResponse 回传工单服务进度 API返回值
// tmall.servicecenter.workcard.serviceprogress.update
//
// 回传工单服务进度
type TmallServicecenterWorkcardServiceprogressUpdateAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardServiceprogressUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardServiceprogressUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterWorkcardServiceprogressUpdateAPIResponseModel).Reset()
}

// TmallServicecenterWorkcardServiceprogressUpdateAPIResponseModel is 回传工单服务进度 成功返回结果
type TmallServicecenterWorkcardServiceprogressUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_serviceprogress_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否调用成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardServiceprogressUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.IsSuccess = false
}

var poolTmallServicecenterWorkcardServiceprogressUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardServiceprogressUpdateAPIResponse)
	},
}

// GetTmallServicecenterWorkcardServiceprogressUpdateAPIResponse 从 sync.Pool 获取 TmallServicecenterWorkcardServiceprogressUpdateAPIResponse
func GetTmallServicecenterWorkcardServiceprogressUpdateAPIResponse() *TmallServicecenterWorkcardServiceprogressUpdateAPIResponse {
	return poolTmallServicecenterWorkcardServiceprogressUpdateAPIResponse.Get().(*TmallServicecenterWorkcardServiceprogressUpdateAPIResponse)
}

// ReleaseTmallServicecenterWorkcardServiceprogressUpdateAPIResponse 将 TmallServicecenterWorkcardServiceprogressUpdateAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterWorkcardServiceprogressUpdateAPIResponse(v *TmallServicecenterWorkcardServiceprogressUpdateAPIResponse) {
	v.Reset()
	poolTmallServicecenterWorkcardServiceprogressUpdateAPIResponse.Put(v)
}
