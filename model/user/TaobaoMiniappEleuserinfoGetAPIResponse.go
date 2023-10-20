package user

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappEleuserinfoGetAPIResponse 获取饿了么用户信息 API返回值
// taobao.miniapp.eleuserinfo.get
//
// 获取饿了么用户信息
type TaobaoMiniappEleuserinfoGetAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappEleuserinfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMiniappEleuserinfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMiniappEleuserinfoGetAPIResponseModel).Reset()
}

// TaobaoMiniappEleuserinfoGetAPIResponseModel is 获取饿了么用户信息 成功返回结果
type TaobaoMiniappEleuserinfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_eleuserinfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// traceId
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 1
	Result *EleUicInfo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMiniappEleuserinfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TraceId = ""
	m.Result = nil
}

var poolTaobaoMiniappEleuserinfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappEleuserinfoGetAPIResponse)
	},
}

// GetTaobaoMiniappEleuserinfoGetAPIResponse 从 sync.Pool 获取 TaobaoMiniappEleuserinfoGetAPIResponse
func GetTaobaoMiniappEleuserinfoGetAPIResponse() *TaobaoMiniappEleuserinfoGetAPIResponse {
	return poolTaobaoMiniappEleuserinfoGetAPIResponse.Get().(*TaobaoMiniappEleuserinfoGetAPIResponse)
}

// ReleaseTaobaoMiniappEleuserinfoGetAPIResponse 将 TaobaoMiniappEleuserinfoGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMiniappEleuserinfoGetAPIResponse(v *TaobaoMiniappEleuserinfoGetAPIResponse) {
	v.Reset()
	poolTaobaoMiniappEleuserinfoGetAPIResponse.Put(v)
}
