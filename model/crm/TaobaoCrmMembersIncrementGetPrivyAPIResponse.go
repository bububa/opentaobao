package crm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmMembersIncrementGetPrivyAPIResponse 增量获取卖家会员 API返回值
// taobao.crm.members.increment.get.privy
//
// 增量获取会员列表，接口返回符合查询条件的所有会员。任何状态更改都会返回,最大允许100
type TaobaoCrmMembersIncrementGetPrivyAPIResponse struct {
	model.CommonResponse
	TaobaoCrmMembersIncrementGetPrivyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCrmMembersIncrementGetPrivyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCrmMembersIncrementGetPrivyAPIResponseModel).Reset()
}

// TaobaoCrmMembersIncrementGetPrivyAPIResponseModel is 增量获取卖家会员 成功返回结果
type TaobaoCrmMembersIncrementGetPrivyAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_members_increment_get_privy_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回当前页的会员列表
	Members []BasicMember `json:"members,omitempty" xml:"members>basic_member,omitempty"`
	// 记录的总条数
	TotalResult int64 `json:"total_result,omitempty" xml:"total_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCrmMembersIncrementGetPrivyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Members = m.Members[:0]
	m.TotalResult = 0
}

var poolTaobaoCrmMembersIncrementGetPrivyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCrmMembersIncrementGetPrivyAPIResponse)
	},
}

// GetTaobaoCrmMembersIncrementGetPrivyAPIResponse 从 sync.Pool 获取 TaobaoCrmMembersIncrementGetPrivyAPIResponse
func GetTaobaoCrmMembersIncrementGetPrivyAPIResponse() *TaobaoCrmMembersIncrementGetPrivyAPIResponse {
	return poolTaobaoCrmMembersIncrementGetPrivyAPIResponse.Get().(*TaobaoCrmMembersIncrementGetPrivyAPIResponse)
}

// ReleaseTaobaoCrmMembersIncrementGetPrivyAPIResponse 将 TaobaoCrmMembersIncrementGetPrivyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCrmMembersIncrementGetPrivyAPIResponse(v *TaobaoCrmMembersIncrementGetPrivyAPIResponse) {
	v.Reset()
	poolTaobaoCrmMembersIncrementGetPrivyAPIResponse.Put(v)
}
