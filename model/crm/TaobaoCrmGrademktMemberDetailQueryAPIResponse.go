package crm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmGrademktMemberDetailQueryAPIResponse 会员等级营销-等级营销活动查询 API返回值
// taobao.crm.grademkt.member.detail.query
//
// 商家通过该接口查询等级营销活动
type TaobaoCrmGrademktMemberDetailQueryAPIResponse struct {
	model.CommonResponse
	TaobaoCrmGrademktMemberDetailQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCrmGrademktMemberDetailQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCrmGrademktMemberDetailQueryAPIResponseModel).Reset()
}

// TaobaoCrmGrademktMemberDetailQueryAPIResponseModel is 会员等级营销-等级营销活动查询 成功返回结果
type TaobaoCrmGrademktMemberDetailQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_grademkt_member_detail_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// totalCount为记录总数
	Model string `json:"model,omitempty" xml:"model,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCrmGrademktMemberDetailQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = ""
}

var poolTaobaoCrmGrademktMemberDetailQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCrmGrademktMemberDetailQueryAPIResponse)
	},
}

// GetTaobaoCrmGrademktMemberDetailQueryAPIResponse 从 sync.Pool 获取 TaobaoCrmGrademktMemberDetailQueryAPIResponse
func GetTaobaoCrmGrademktMemberDetailQueryAPIResponse() *TaobaoCrmGrademktMemberDetailQueryAPIResponse {
	return poolTaobaoCrmGrademktMemberDetailQueryAPIResponse.Get().(*TaobaoCrmGrademktMemberDetailQueryAPIResponse)
}

// ReleaseTaobaoCrmGrademktMemberDetailQueryAPIResponse 将 TaobaoCrmGrademktMemberDetailQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCrmGrademktMemberDetailQueryAPIResponse(v *TaobaoCrmGrademktMemberDetailQueryAPIResponse) {
	v.Reset()
	poolTaobaoCrmGrademktMemberDetailQueryAPIResponse.Put(v)
}
