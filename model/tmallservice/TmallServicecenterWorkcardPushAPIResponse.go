package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardPushAPIResponse 推送服务工单信息 API返回值
// tmall.servicecenter.workcard.push
//
// 服务商家推送工单信息到天猫。
type TmallServicecenterWorkcardPushAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardPushAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardPushAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterWorkcardPushAPIResponseModel).Reset()
}

// TmallServicecenterWorkcardPushAPIResponseModel is 推送服务工单信息 成功返回结果
type TmallServicecenterWorkcardPushAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_push_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	WorkcardResult *ResultBase `json:"workcard_result,omitempty" xml:"workcard_result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardPushAPIResponseModel) Reset() {
	m.RequestId = ""
	m.WorkcardResult = nil
}

var poolTmallServicecenterWorkcardPushAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardPushAPIResponse)
	},
}

// GetTmallServicecenterWorkcardPushAPIResponse 从 sync.Pool 获取 TmallServicecenterWorkcardPushAPIResponse
func GetTmallServicecenterWorkcardPushAPIResponse() *TmallServicecenterWorkcardPushAPIResponse {
	return poolTmallServicecenterWorkcardPushAPIResponse.Get().(*TmallServicecenterWorkcardPushAPIResponse)
}

// ReleaseTmallServicecenterWorkcardPushAPIResponse 将 TmallServicecenterWorkcardPushAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterWorkcardPushAPIResponse(v *TmallServicecenterWorkcardPushAPIResponse) {
	v.Reset()
	poolTmallServicecenterWorkcardPushAPIResponse.Put(v)
}
