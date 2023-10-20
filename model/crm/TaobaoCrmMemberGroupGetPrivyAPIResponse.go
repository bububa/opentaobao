package crm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmMemberGroupGetPrivyAPIResponse 获取买家身上的标签（隐私号版） API返回值
// taobao.crm.member.group.get.privy
//
// 获取买家身上的标签，不返回标签的总人数
type TaobaoCrmMemberGroupGetPrivyAPIResponse struct {
	model.CommonResponse
	TaobaoCrmMemberGroupGetPrivyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCrmMemberGroupGetPrivyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCrmMemberGroupGetPrivyAPIResponseModel).Reset()
}

// TaobaoCrmMemberGroupGetPrivyAPIResponseModel is 获取买家身上的标签（隐私号版） 成功返回结果
type TaobaoCrmMemberGroupGetPrivyAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_member_group_get_privy_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询到的当前卖家的当前页的会员
	Groups []Group `json:"groups,omitempty" xml:"groups>group,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCrmMemberGroupGetPrivyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Groups = m.Groups[:0]
}

var poolTaobaoCrmMemberGroupGetPrivyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCrmMemberGroupGetPrivyAPIResponse)
	},
}

// GetTaobaoCrmMemberGroupGetPrivyAPIResponse 从 sync.Pool 获取 TaobaoCrmMemberGroupGetPrivyAPIResponse
func GetTaobaoCrmMemberGroupGetPrivyAPIResponse() *TaobaoCrmMemberGroupGetPrivyAPIResponse {
	return poolTaobaoCrmMemberGroupGetPrivyAPIResponse.Get().(*TaobaoCrmMemberGroupGetPrivyAPIResponse)
}

// ReleaseTaobaoCrmMemberGroupGetPrivyAPIResponse 将 TaobaoCrmMemberGroupGetPrivyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCrmMemberGroupGetPrivyAPIResponse(v *TaobaoCrmMemberGroupGetPrivyAPIResponse) {
	v.Reset()
	poolTaobaoCrmMemberGroupGetPrivyAPIResponse.Put(v)
}
