package game

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAppleNewuserSignNotifyAPIResponse 新用户签约通知接口 API返回值
// taobao.apple.newuser.sign.notify
//
// 用户付款成功后，资和信主动通知签约结果
type TaobaoAppleNewuserSignNotifyAPIResponse struct {
	model.CommonResponse
	TaobaoAppleNewuserSignNotifyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAppleNewuserSignNotifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAppleNewuserSignNotifyAPIResponseModel).Reset()
}

// TaobaoAppleNewuserSignNotifyAPIResponseModel is 新用户签约通知接口 成功返回结果
type TaobaoAppleNewuserSignNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"apple_newuser_sign_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 处理结果说明
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 处理结果码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAppleNewuserSignNotifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.ResultCode = ""
}

var poolTaobaoAppleNewuserSignNotifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAppleNewuserSignNotifyAPIResponse)
	},
}

// GetTaobaoAppleNewuserSignNotifyAPIResponse 从 sync.Pool 获取 TaobaoAppleNewuserSignNotifyAPIResponse
func GetTaobaoAppleNewuserSignNotifyAPIResponse() *TaobaoAppleNewuserSignNotifyAPIResponse {
	return poolTaobaoAppleNewuserSignNotifyAPIResponse.Get().(*TaobaoAppleNewuserSignNotifyAPIResponse)
}

// ReleaseTaobaoAppleNewuserSignNotifyAPIResponse 将 TaobaoAppleNewuserSignNotifyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAppleNewuserSignNotifyAPIResponse(v *TaobaoAppleNewuserSignNotifyAPIResponse) {
	v.Reset()
	poolTaobaoAppleNewuserSignNotifyAPIResponse.Put(v)
}
