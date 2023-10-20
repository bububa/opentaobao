package game

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAppleNewuserSignNotifyNewversionAPIResponse 新用户签约结果通知接口v2 API返回值
// taobao.apple.newuser.sign.notify.newversion
//
// 资和信主动通知签约结果
type TaobaoAppleNewuserSignNotifyNewversionAPIResponse struct {
	model.CommonResponse
	TaobaoAppleNewuserSignNotifyNewversionAPIResponseModel
}

// TaobaoAppleNewuserSignNotifyNewversionAPIResponseModel is 新用户签约结果通知接口v2 成功返回结果
type TaobaoAppleNewuserSignNotifyNewversionAPIResponseModel struct {
	XMLName xml.Name `xml:"apple_newuser_sign_notify_newversion_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 处理结果说明
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 处理结果码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
}
