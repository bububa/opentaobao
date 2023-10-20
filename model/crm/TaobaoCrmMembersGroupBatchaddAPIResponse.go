package crm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmMembersGroupBatchaddAPIResponse 给一批会员添加一个分组 API返回值
// taobao.crm.members.group.batchadd
//
// 为一批会员添加分组，接口返回添加是否成功,如至少有一个会员的分组添加成功，接口就返回成功，否则返回失败，如果当前会员已经拥有当前分组，则直接跳过
type TaobaoCrmMembersGroupBatchaddAPIResponse struct {
	model.CommonResponse
	TaobaoCrmMembersGroupBatchaddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCrmMembersGroupBatchaddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCrmMembersGroupBatchaddAPIResponseModel).Reset()
}

// TaobaoCrmMembersGroupBatchaddAPIResponseModel is 给一批会员添加一个分组 成功返回结果
type TaobaoCrmMembersGroupBatchaddAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_members_group_batchadd_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 添加操作是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCrmMembersGroupBatchaddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoCrmMembersGroupBatchaddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCrmMembersGroupBatchaddAPIResponse)
	},
}

// GetTaobaoCrmMembersGroupBatchaddAPIResponse 从 sync.Pool 获取 TaobaoCrmMembersGroupBatchaddAPIResponse
func GetTaobaoCrmMembersGroupBatchaddAPIResponse() *TaobaoCrmMembersGroupBatchaddAPIResponse {
	return poolTaobaoCrmMembersGroupBatchaddAPIResponse.Get().(*TaobaoCrmMembersGroupBatchaddAPIResponse)
}

// ReleaseTaobaoCrmMembersGroupBatchaddAPIResponse 将 TaobaoCrmMembersGroupBatchaddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCrmMembersGroupBatchaddAPIResponse(v *TaobaoCrmMembersGroupBatchaddAPIResponse) {
	v.Reset()
	poolTaobaoCrmMembersGroupBatchaddAPIResponse.Put(v)
}
