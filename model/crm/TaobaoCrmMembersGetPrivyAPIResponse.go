package crm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmMembersGetPrivyAPIResponse 获取卖家会员(基本查询) API返回值
// taobao.crm.members.get.privy
//
// 查询卖家的会员，进行基本的查询，返回符合条件的会员列表
type TaobaoCrmMembersGetPrivyAPIResponse struct {
	model.CommonResponse
	TaobaoCrmMembersGetPrivyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCrmMembersGetPrivyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCrmMembersGetPrivyAPIResponseModel).Reset()
}

// TaobaoCrmMembersGetPrivyAPIResponseModel is 获取卖家会员(基本查询) 成功返回结果
type TaobaoCrmMembersGetPrivyAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_members_get_privy_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 根据一定条件查询到卖家的会员
	Members []BasicMember `json:"members,omitempty" xml:"members>basic_member,omitempty"`
	// 记录总数
	TotalResult int64 `json:"total_result,omitempty" xml:"total_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCrmMembersGetPrivyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Members = m.Members[:0]
	m.TotalResult = 0
}

var poolTaobaoCrmMembersGetPrivyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCrmMembersGetPrivyAPIResponse)
	},
}

// GetTaobaoCrmMembersGetPrivyAPIResponse 从 sync.Pool 获取 TaobaoCrmMembersGetPrivyAPIResponse
func GetTaobaoCrmMembersGetPrivyAPIResponse() *TaobaoCrmMembersGetPrivyAPIResponse {
	return poolTaobaoCrmMembersGetPrivyAPIResponse.Get().(*TaobaoCrmMembersGetPrivyAPIResponse)
}

// ReleaseTaobaoCrmMembersGetPrivyAPIResponse 将 TaobaoCrmMembersGetPrivyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCrmMembersGetPrivyAPIResponse(v *TaobaoCrmMembersGetPrivyAPIResponse) {
	v.Reset()
	poolTaobaoCrmMembersGetPrivyAPIResponse.Put(v)
}
