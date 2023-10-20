package game

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAppleCardActiveApplyNotifyAPIResponse 苹果卡密申请激活回调接口 API返回值
// taobao.apple.card.active.apply.notify
//
// 苹果卡密申请激活回调接口
type TaobaoAppleCardActiveApplyNotifyAPIResponse struct {
	model.CommonResponse
	TaobaoAppleCardActiveApplyNotifyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAppleCardActiveApplyNotifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAppleCardActiveApplyNotifyAPIResponseModel).Reset()
}

// TaobaoAppleCardActiveApplyNotifyAPIResponseModel is 苹果卡密申请激活回调接口 成功返回结果
type TaobaoAppleCardActiveApplyNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"apple_card_active_apply_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAppleCardActiveApplyNotifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = ""
	m.ResultMsg = ""
}

var poolTaobaoAppleCardActiveApplyNotifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAppleCardActiveApplyNotifyAPIResponse)
	},
}

// GetTaobaoAppleCardActiveApplyNotifyAPIResponse 从 sync.Pool 获取 TaobaoAppleCardActiveApplyNotifyAPIResponse
func GetTaobaoAppleCardActiveApplyNotifyAPIResponse() *TaobaoAppleCardActiveApplyNotifyAPIResponse {
	return poolTaobaoAppleCardActiveApplyNotifyAPIResponse.Get().(*TaobaoAppleCardActiveApplyNotifyAPIResponse)
}

// ReleaseTaobaoAppleCardActiveApplyNotifyAPIResponse 将 TaobaoAppleCardActiveApplyNotifyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAppleCardActiveApplyNotifyAPIResponse(v *TaobaoAppleCardActiveApplyNotifyAPIResponse) {
	v.Reset()
	poolTaobaoAppleCardActiveApplyNotifyAPIResponse.Put(v)
}
