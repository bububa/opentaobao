package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScInvitecodeGetAPIResponse 淘宝客-公用-私域用户邀请码生成 API返回值
// taobao.tbk.sc.invitecode.get
//
// 私域用户管理(即渠道管理或会员运营管理)功能中，通过此API可生成淘宝客自身的邀请码。
type TaobaoTbkScInvitecodeGetAPIResponse struct {
	model.CommonResponse
	TaobaoTbkScInvitecodeGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkScInvitecodeGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkScInvitecodeGetAPIResponseModel).Reset()
}

// TaobaoTbkScInvitecodeGetAPIResponseModel is 淘宝客-公用-私域用户邀请码生成 成功返回结果
type TaobaoTbkScInvitecodeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_invitecode_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data *TaobaoTbkScInvitecodeGetData `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkScInvitecodeGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoTbkScInvitecodeGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScInvitecodeGetAPIResponse)
	},
}

// GetTaobaoTbkScInvitecodeGetAPIResponse 从 sync.Pool 获取 TaobaoTbkScInvitecodeGetAPIResponse
func GetTaobaoTbkScInvitecodeGetAPIResponse() *TaobaoTbkScInvitecodeGetAPIResponse {
	return poolTaobaoTbkScInvitecodeGetAPIResponse.Get().(*TaobaoTbkScInvitecodeGetAPIResponse)
}

// ReleaseTaobaoTbkScInvitecodeGetAPIResponse 将 TaobaoTbkScInvitecodeGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkScInvitecodeGetAPIResponse(v *TaobaoTbkScInvitecodeGetAPIResponse) {
	v.Reset()
	poolTaobaoTbkScInvitecodeGetAPIResponse.Put(v)
}
