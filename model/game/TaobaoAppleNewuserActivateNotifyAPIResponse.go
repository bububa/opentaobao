package game

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAppleNewuserActivateNotifyAPIResponse 新用户激活通知接口 API返回值
// taobao.apple.newuser.activate.notify
//
// 资和信主动通知激活结果
type TaobaoAppleNewuserActivateNotifyAPIResponse struct {
	model.CommonResponse
	TaobaoAppleNewuserActivateNotifyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAppleNewuserActivateNotifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAppleNewuserActivateNotifyAPIResponseModel).Reset()
}

// TaobaoAppleNewuserActivateNotifyAPIResponseModel is 新用户激活通知接口 成功返回结果
type TaobaoAppleNewuserActivateNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"apple_newuser_activate_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 处理结果说明
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 处理结果码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAppleNewuserActivateNotifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.ResultCode = ""
}

var poolTaobaoAppleNewuserActivateNotifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAppleNewuserActivateNotifyAPIResponse)
	},
}

// GetTaobaoAppleNewuserActivateNotifyAPIResponse 从 sync.Pool 获取 TaobaoAppleNewuserActivateNotifyAPIResponse
func GetTaobaoAppleNewuserActivateNotifyAPIResponse() *TaobaoAppleNewuserActivateNotifyAPIResponse {
	return poolTaobaoAppleNewuserActivateNotifyAPIResponse.Get().(*TaobaoAppleNewuserActivateNotifyAPIResponse)
}

// ReleaseTaobaoAppleNewuserActivateNotifyAPIResponse 将 TaobaoAppleNewuserActivateNotifyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAppleNewuserActivateNotifyAPIResponse(v *TaobaoAppleNewuserActivateNotifyAPIResponse) {
	v.Reset()
	poolTaobaoAppleNewuserActivateNotifyAPIResponse.Put(v)
}
