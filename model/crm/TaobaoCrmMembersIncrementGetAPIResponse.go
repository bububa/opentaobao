package crm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmMembersIncrementGetAPIResponse 增量获取卖家会员 API返回值
// taobao.crm.members.increment.get
//
// 增量获取会员列表，接口返回符合查询条件的所有会员。任何状态更改都会返回,最大允许100
type TaobaoCrmMembersIncrementGetAPIResponse struct {
	model.CommonResponse
	TaobaoCrmMembersIncrementGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCrmMembersIncrementGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCrmMembersIncrementGetAPIResponseModel).Reset()
}

// TaobaoCrmMembersIncrementGetAPIResponseModel is 增量获取卖家会员 成功返回结果
type TaobaoCrmMembersIncrementGetAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_members_increment_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回当前页的会员列表
	Members []BasicMember `json:"members,omitempty" xml:"members>basic_member,omitempty"`
	// 记录的总条数
	TotalResult int64 `json:"total_result,omitempty" xml:"total_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCrmMembersIncrementGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Members = m.Members[:0]
	m.TotalResult = 0
}

var poolTaobaoCrmMembersIncrementGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCrmMembersIncrementGetAPIResponse)
	},
}

// GetTaobaoCrmMembersIncrementGetAPIResponse 从 sync.Pool 获取 TaobaoCrmMembersIncrementGetAPIResponse
func GetTaobaoCrmMembersIncrementGetAPIResponse() *TaobaoCrmMembersIncrementGetAPIResponse {
	return poolTaobaoCrmMembersIncrementGetAPIResponse.Get().(*TaobaoCrmMembersIncrementGetAPIResponse)
}

// ReleaseTaobaoCrmMembersIncrementGetAPIResponse 将 TaobaoCrmMembersIncrementGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCrmMembersIncrementGetAPIResponse(v *TaobaoCrmMembersIncrementGetAPIResponse) {
	v.Reset()
	poolTaobaoCrmMembersIncrementGetAPIResponse.Put(v)
}
