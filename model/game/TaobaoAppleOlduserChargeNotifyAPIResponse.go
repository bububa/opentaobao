package game

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAppleOlduserChargeNotifyAPIResponse 老用户激活并兑换通知接口 API返回值
// taobao.apple.olduser.charge.notify
//
// 老用户激活并兑换通知接口
type TaobaoAppleOlduserChargeNotifyAPIResponse struct {
	model.CommonResponse
	TaobaoAppleOlduserChargeNotifyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAppleOlduserChargeNotifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAppleOlduserChargeNotifyAPIResponseModel).Reset()
}

// TaobaoAppleOlduserChargeNotifyAPIResponseModel is 老用户激活并兑换通知接口 成功返回结果
type TaobaoAppleOlduserChargeNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"apple_olduser_charge_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 处理结果说明
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 处理结果码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAppleOlduserChargeNotifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.ResultCode = ""
}

var poolTaobaoAppleOlduserChargeNotifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAppleOlduserChargeNotifyAPIResponse)
	},
}

// GetTaobaoAppleOlduserChargeNotifyAPIResponse 从 sync.Pool 获取 TaobaoAppleOlduserChargeNotifyAPIResponse
func GetTaobaoAppleOlduserChargeNotifyAPIResponse() *TaobaoAppleOlduserChargeNotifyAPIResponse {
	return poolTaobaoAppleOlduserChargeNotifyAPIResponse.Get().(*TaobaoAppleOlduserChargeNotifyAPIResponse)
}

// ReleaseTaobaoAppleOlduserChargeNotifyAPIResponse 将 TaobaoAppleOlduserChargeNotifyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAppleOlduserChargeNotifyAPIResponse(v *TaobaoAppleOlduserChargeNotifyAPIResponse) {
	v.Reset()
	poolTaobaoAppleOlduserChargeNotifyAPIResponse.Put(v)
}
