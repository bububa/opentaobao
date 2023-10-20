package crm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmMembersGetAPIResponse 获取卖家的会员（基本查询） API返回值
// taobao.crm.members.get
//
// 查询卖家的会员，进行基本的查询，返回符合条件的会员列表
type TaobaoCrmMembersGetAPIResponse struct {
	model.CommonResponse
	TaobaoCrmMembersGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCrmMembersGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCrmMembersGetAPIResponseModel).Reset()
}

// TaobaoCrmMembersGetAPIResponseModel is 获取卖家的会员（基本查询） 成功返回结果
type TaobaoCrmMembersGetAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_members_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 根据一定条件查询到卖家的会员
	Members []BasicMember `json:"members,omitempty" xml:"members>basic_member,omitempty"`
	// 记录总数
	TotalResult int64 `json:"total_result,omitempty" xml:"total_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCrmMembersGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Members = m.Members[:0]
	m.TotalResult = 0
}

var poolTaobaoCrmMembersGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCrmMembersGetAPIResponse)
	},
}

// GetTaobaoCrmMembersGetAPIResponse 从 sync.Pool 获取 TaobaoCrmMembersGetAPIResponse
func GetTaobaoCrmMembersGetAPIResponse() *TaobaoCrmMembersGetAPIResponse {
	return poolTaobaoCrmMembersGetAPIResponse.Get().(*TaobaoCrmMembersGetAPIResponse)
}

// ReleaseTaobaoCrmMembersGetAPIResponse 将 TaobaoCrmMembersGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCrmMembersGetAPIResponse(v *TaobaoCrmMembersGetAPIResponse) {
	v.Reset()
	poolTaobaoCrmMembersGetAPIResponse.Put(v)
}
