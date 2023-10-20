package tmallfcbox

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallFcboxNotifyAPIResponse 丰巢通知接口 API返回值
// tmall.fcbox.notify
//
// tmax接收丰巢快递通知
type TmallFcboxNotifyAPIResponse struct {
	model.CommonResponse
	TmallFcboxNotifyAPIResponseModel
}

// Reset 清空结构体
func (m *TmallFcboxNotifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallFcboxNotifyAPIResponseModel).Reset()
}

// TmallFcboxNotifyAPIResponseModel is 丰巢通知接口 成功返回结果
type TmallFcboxNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_fcbox_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallFcboxNotifyResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallFcboxNotifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallFcboxNotifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallFcboxNotifyAPIResponse)
	},
}

// GetTmallFcboxNotifyAPIResponse 从 sync.Pool 获取 TmallFcboxNotifyAPIResponse
func GetTmallFcboxNotifyAPIResponse() *TmallFcboxNotifyAPIResponse {
	return poolTmallFcboxNotifyAPIResponse.Get().(*TmallFcboxNotifyAPIResponse)
}

// ReleaseTmallFcboxNotifyAPIResponse 将 TmallFcboxNotifyAPIResponse 保存到 sync.Pool
func ReleaseTmallFcboxNotifyAPIResponse(v *TmallFcboxNotifyAPIResponse) {
	v.Reset()
	poolTmallFcboxNotifyAPIResponse.Put(v)
}
