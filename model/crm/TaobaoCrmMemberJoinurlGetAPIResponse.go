package crm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmMemberJoinurlGetAPIResponse 会员入会地址获取 API返回值
// taobao.crm.member.joinurl.get
//
// 会员入会地址获取
type TaobaoCrmMemberJoinurlGetAPIResponse struct {
	model.CommonResponse
	TaobaoCrmMemberJoinurlGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCrmMemberJoinurlGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCrmMemberJoinurlGetAPIResponseModel).Reset()
}

// TaobaoCrmMemberJoinurlGetAPIResponseModel is 会员入会地址获取 成功返回结果
type TaobaoCrmMemberJoinurlGetAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_member_joinurl_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoCrmMemberJoinurlGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCrmMemberJoinurlGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoCrmMemberJoinurlGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCrmMemberJoinurlGetAPIResponse)
	},
}

// GetTaobaoCrmMemberJoinurlGetAPIResponse 从 sync.Pool 获取 TaobaoCrmMemberJoinurlGetAPIResponse
func GetTaobaoCrmMemberJoinurlGetAPIResponse() *TaobaoCrmMemberJoinurlGetAPIResponse {
	return poolTaobaoCrmMemberJoinurlGetAPIResponse.Get().(*TaobaoCrmMemberJoinurlGetAPIResponse)
}

// ReleaseTaobaoCrmMemberJoinurlGetAPIResponse 将 TaobaoCrmMemberJoinurlGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCrmMemberJoinurlGetAPIResponse(v *TaobaoCrmMemberJoinurlGetAPIResponse) {
	v.Reset()
	poolTaobaoCrmMemberJoinurlGetAPIResponse.Put(v)
}
