package tmallsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardDeliveryAPIResponse 开始配送工单 API返回值
// tmall.servicecenter.workcard.delivery
//
// 服务商调用该接口通知天猫服务平台服务商工人已开始配送工单
type TmallServicecenterWorkcardDeliveryAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardDeliveryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardDeliveryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterWorkcardDeliveryAPIResponseModel).Reset()
}

// TmallServicecenterWorkcardDeliveryAPIResponseModel is 开始配送工单 成功返回结果
type TmallServicecenterWorkcardDeliveryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_delivery_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回信息
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回code
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	MsgSuccess bool `json:"msg_success,omitempty" xml:"msg_success,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardDeliveryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.MsgSuccess = false
}

var poolTmallServicecenterWorkcardDeliveryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardDeliveryAPIResponse)
	},
}

// GetTmallServicecenterWorkcardDeliveryAPIResponse 从 sync.Pool 获取 TmallServicecenterWorkcardDeliveryAPIResponse
func GetTmallServicecenterWorkcardDeliveryAPIResponse() *TmallServicecenterWorkcardDeliveryAPIResponse {
	return poolTmallServicecenterWorkcardDeliveryAPIResponse.Get().(*TmallServicecenterWorkcardDeliveryAPIResponse)
}

// ReleaseTmallServicecenterWorkcardDeliveryAPIResponse 将 TmallServicecenterWorkcardDeliveryAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterWorkcardDeliveryAPIResponse(v *TmallServicecenterWorkcardDeliveryAPIResponse) {
	v.Reset()
	poolTmallServicecenterWorkcardDeliveryAPIResponse.Put(v)
}
