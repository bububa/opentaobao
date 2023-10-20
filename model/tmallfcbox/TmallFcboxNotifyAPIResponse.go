package tmallfcbox

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallfcboxnotifyAPIResponse 丰巢通知接口 API返回值
// tmall.fcbox.notify
//
// tmax接收丰巢快递通知
type TmallfcboxnotifyAPIResponse struct {
	model.CommonResponse
	TmallfcboxnotifyAPIResponseModel
}

// TmallfcboxnotifyAPIResponseModel is 丰巢通知接口 成功返回结果
type TmallfcboxnotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_fcbox_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallfcboxnotifyResult `json:"result,omitempty" xml:"result,omitempty"`
}
