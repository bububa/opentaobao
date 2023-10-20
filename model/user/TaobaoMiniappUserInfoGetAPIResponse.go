package user

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappUserInfoGetAPIResponse 用户开放信息获取 API返回值
// taobao.miniapp.userInfo.get
//
// 获取用户的 openId，snsNick（如果用户设置过的话），和加密头像链接
type TaobaoMiniappUserInfoGetAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappUserInfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMiniappUserInfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMiniappUserInfoGetAPIResponseModel).Reset()
}

// TaobaoMiniappUserInfoGetAPIResponseModel is 用户开放信息获取 成功返回结果
type TaobaoMiniappUserInfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_userInfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoMiniappUserInfoGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMiniappUserInfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoMiniappUserInfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappUserInfoGetAPIResponse)
	},
}

// GetTaobaoMiniappUserInfoGetAPIResponse 从 sync.Pool 获取 TaobaoMiniappUserInfoGetAPIResponse
func GetTaobaoMiniappUserInfoGetAPIResponse() *TaobaoMiniappUserInfoGetAPIResponse {
	return poolTaobaoMiniappUserInfoGetAPIResponse.Get().(*TaobaoMiniappUserInfoGetAPIResponse)
}

// ReleaseTaobaoMiniappUserInfoGetAPIResponse 将 TaobaoMiniappUserInfoGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMiniappUserInfoGetAPIResponse(v *TaobaoMiniappUserInfoGetAPIResponse) {
	v.Reset()
	poolTaobaoMiniappUserInfoGetAPIResponse.Put(v)
}
