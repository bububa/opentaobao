package crm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmMemberIdentityGetAPIResponse 会员身份识别 API返回值
// taobao.crm.member.identity.get
//
// 用来识别该用户是否是商家会员
type TaobaoCrmMemberIdentityGetAPIResponse struct {
	model.CommonResponse
	TaobaoCrmMemberIdentityGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCrmMemberIdentityGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCrmMemberIdentityGetAPIResponseModel).Reset()
}

// TaobaoCrmMemberIdentityGetAPIResponseModel is 会员身份识别 成功返回结果
type TaobaoCrmMemberIdentityGetAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_member_identity_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoCrmMemberIdentityGetResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCrmMemberIdentityGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoCrmMemberIdentityGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCrmMemberIdentityGetAPIResponse)
	},
}

// GetTaobaoCrmMemberIdentityGetAPIResponse 从 sync.Pool 获取 TaobaoCrmMemberIdentityGetAPIResponse
func GetTaobaoCrmMemberIdentityGetAPIResponse() *TaobaoCrmMemberIdentityGetAPIResponse {
	return poolTaobaoCrmMemberIdentityGetAPIResponse.Get().(*TaobaoCrmMemberIdentityGetAPIResponse)
}

// ReleaseTaobaoCrmMemberIdentityGetAPIResponse 将 TaobaoCrmMemberIdentityGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCrmMemberIdentityGetAPIResponse(v *TaobaoCrmMemberIdentityGetAPIResponse) {
	v.Reset()
	poolTaobaoCrmMemberIdentityGetAPIResponse.Put(v)
}
