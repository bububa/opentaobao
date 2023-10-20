package user

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappEleuserPhoneGetAPIResponse 获取饿了么用户信息 API返回值
// taobao.miniapp.eleuser.phone.get
//
// 获取饿了么用户信息
type TaobaoMiniappEleuserPhoneGetAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappEleuserPhoneGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMiniappEleuserPhoneGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMiniappEleuserPhoneGetAPIResponseModel).Reset()
}

// TaobaoMiniappEleuserPhoneGetAPIResponseModel is 获取饿了么用户信息 成功返回结果
type TaobaoMiniappEleuserPhoneGetAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_eleuser_phone_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// traceId
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 返回对象
	Result *EleUicInfo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMiniappEleuserPhoneGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TraceId = ""
	m.Result = nil
}

var poolTaobaoMiniappEleuserPhoneGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappEleuserPhoneGetAPIResponse)
	},
}

// GetTaobaoMiniappEleuserPhoneGetAPIResponse 从 sync.Pool 获取 TaobaoMiniappEleuserPhoneGetAPIResponse
func GetTaobaoMiniappEleuserPhoneGetAPIResponse() *TaobaoMiniappEleuserPhoneGetAPIResponse {
	return poolTaobaoMiniappEleuserPhoneGetAPIResponse.Get().(*TaobaoMiniappEleuserPhoneGetAPIResponse)
}

// ReleaseTaobaoMiniappEleuserPhoneGetAPIResponse 将 TaobaoMiniappEleuserPhoneGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMiniappEleuserPhoneGetAPIResponse(v *TaobaoMiniappEleuserPhoneGetAPIResponse) {
	v.Reset()
	poolTaobaoMiniappEleuserPhoneGetAPIResponse.Put(v)
}
