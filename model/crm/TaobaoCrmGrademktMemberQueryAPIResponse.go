package crm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmGrademktMemberQueryAPIResponse 会员等级营销-会员关系查询 API返回值
// taobao.crm.grademkt.member.query
//
// 商家通过该接口查询线上店铺会员。
type TaobaoCrmGrademktMemberQueryAPIResponse struct {
	model.CommonResponse
	TaobaoCrmGrademktMemberQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCrmGrademktMemberQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCrmGrademktMemberQueryAPIResponseModel).Reset()
}

// TaobaoCrmGrademktMemberQueryAPIResponseModel is 会员等级营销-会员关系查询 成功返回结果
type TaobaoCrmGrademktMemberQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_grademkt_member_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// json格式
	Module string `json:"module,omitempty" xml:"module,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCrmGrademktMemberQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Module = ""
}

var poolTaobaoCrmGrademktMemberQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCrmGrademktMemberQueryAPIResponse)
	},
}

// GetTaobaoCrmGrademktMemberQueryAPIResponse 从 sync.Pool 获取 TaobaoCrmGrademktMemberQueryAPIResponse
func GetTaobaoCrmGrademktMemberQueryAPIResponse() *TaobaoCrmGrademktMemberQueryAPIResponse {
	return poolTaobaoCrmGrademktMemberQueryAPIResponse.Get().(*TaobaoCrmGrademktMemberQueryAPIResponse)
}

// ReleaseTaobaoCrmGrademktMemberQueryAPIResponse 将 TaobaoCrmGrademktMemberQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCrmGrademktMemberQueryAPIResponse(v *TaobaoCrmGrademktMemberQueryAPIResponse) {
	v.Reset()
	poolTaobaoCrmGrademktMemberQueryAPIResponse.Put(v)
}
