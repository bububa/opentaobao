package crm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmGroupsGetAPIResponse 查询卖家的分组 API返回值
// taobao.crm.groups.get
//
// 查询卖家的分组，返回查询到的分组列表，分页返回分组
type TaobaoCrmGroupsGetAPIResponse struct {
	model.CommonResponse
	TaobaoCrmGroupsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCrmGroupsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCrmGroupsGetAPIResponseModel).Reset()
}

// TaobaoCrmGroupsGetAPIResponseModel is 查询卖家的分组 成功返回结果
type TaobaoCrmGroupsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_groups_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询到的当前卖家的当前页的会员
	Groups []Group `json:"groups,omitempty" xml:"groups>group,omitempty"`
	// 记录总数
	TotalResult int64 `json:"total_result,omitempty" xml:"total_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCrmGroupsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Groups = m.Groups[:0]
	m.TotalResult = 0
}

var poolTaobaoCrmGroupsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCrmGroupsGetAPIResponse)
	},
}

// GetTaobaoCrmGroupsGetAPIResponse 从 sync.Pool 获取 TaobaoCrmGroupsGetAPIResponse
func GetTaobaoCrmGroupsGetAPIResponse() *TaobaoCrmGroupsGetAPIResponse {
	return poolTaobaoCrmGroupsGetAPIResponse.Get().(*TaobaoCrmGroupsGetAPIResponse)
}

// ReleaseTaobaoCrmGroupsGetAPIResponse 将 TaobaoCrmGroupsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCrmGroupsGetAPIResponse(v *TaobaoCrmGroupsGetAPIResponse) {
	v.Reset()
	poolTaobaoCrmGroupsGetAPIResponse.Put(v)
}
