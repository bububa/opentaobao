package subuser

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubusersGetAPIResponse 获取指定账户的子账号简易信息列表 API返回值
// taobao.subusers.get
//
// 获取主账号下的子账号简易账号信息集合。（只能通过主账号登陆并且查询该属于主账号的子账号信息）
type TaobaoSubusersGetAPIResponse struct {
	model.CommonResponse
	TaobaoSubusersGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSubusersGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSubusersGetAPIResponseModel).Reset()
}

// TaobaoSubusersGetAPIResponseModel is 获取指定账户的子账号简易信息列表 成功返回结果
type TaobaoSubusersGetAPIResponseModel struct {
	XMLName xml.Name `xml:"subusers_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 子账号基本信息
	Subaccounts []SubAccountInfo `json:"subaccounts,omitempty" xml:"subaccounts>sub_account_info,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSubusersGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Subaccounts = m.Subaccounts[:0]
}

var poolTaobaoSubusersGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSubusersGetAPIResponse)
	},
}

// GetTaobaoSubusersGetAPIResponse 从 sync.Pool 获取 TaobaoSubusersGetAPIResponse
func GetTaobaoSubusersGetAPIResponse() *TaobaoSubusersGetAPIResponse {
	return poolTaobaoSubusersGetAPIResponse.Get().(*TaobaoSubusersGetAPIResponse)
}

// ReleaseTaobaoSubusersGetAPIResponse 将 TaobaoSubusersGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSubusersGetAPIResponse(v *TaobaoSubusersGetAPIResponse) {
	v.Reset()
	poolTaobaoSubusersGetAPIResponse.Put(v)
}
