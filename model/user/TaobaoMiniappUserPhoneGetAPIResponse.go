package user

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappUserPhoneGetAPIResponse 获取当前授权用户手机号码 API返回值
// taobao.miniapp.user.phone.get
//
// 在商家应用中，获取当前授权用户手机号码
type TaobaoMiniappUserPhoneGetAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappUserPhoneGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMiniappUserPhoneGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMiniappUserPhoneGetAPIResponseModel).Reset()
}

// TaobaoMiniappUserPhoneGetAPIResponseModel is 获取当前授权用户手机号码 成功返回结果
type TaobaoMiniappUserPhoneGetAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_user_phone_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 用户手机号码
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMiniappUserPhoneGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Phone = ""
}

var poolTaobaoMiniappUserPhoneGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappUserPhoneGetAPIResponse)
	},
}

// GetTaobaoMiniappUserPhoneGetAPIResponse 从 sync.Pool 获取 TaobaoMiniappUserPhoneGetAPIResponse
func GetTaobaoMiniappUserPhoneGetAPIResponse() *TaobaoMiniappUserPhoneGetAPIResponse {
	return poolTaobaoMiniappUserPhoneGetAPIResponse.Get().(*TaobaoMiniappUserPhoneGetAPIResponse)
}

// ReleaseTaobaoMiniappUserPhoneGetAPIResponse 将 TaobaoMiniappUserPhoneGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMiniappUserPhoneGetAPIResponse(v *TaobaoMiniappUserPhoneGetAPIResponse) {
	v.Reset()
	poolTaobaoMiniappUserPhoneGetAPIResponse.Put(v)
}
