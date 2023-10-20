package crm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmGrademktMemberDetailDeleteAPIResponse 会员等级营销-删除商品等级营销明细 API返回值
// taobao.crm.grademkt.member.detail.delete
//
// 删除商品等级营销明细
type TaobaoCrmGrademktMemberDetailDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoCrmGrademktMemberDetailDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCrmGrademktMemberDetailDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCrmGrademktMemberDetailDeleteAPIResponseModel).Reset()
}

// TaobaoCrmGrademktMemberDetailDeleteAPIResponseModel is 会员等级营销-删除商品等级营销明细 成功返回结果
type TaobaoCrmGrademktMemberDetailDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_grademkt_member_detail_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作是否成功
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCrmGrademktMemberDetailDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Module = false
}

var poolTaobaoCrmGrademktMemberDetailDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCrmGrademktMemberDetailDeleteAPIResponse)
	},
}

// GetTaobaoCrmGrademktMemberDetailDeleteAPIResponse 从 sync.Pool 获取 TaobaoCrmGrademktMemberDetailDeleteAPIResponse
func GetTaobaoCrmGrademktMemberDetailDeleteAPIResponse() *TaobaoCrmGrademktMemberDetailDeleteAPIResponse {
	return poolTaobaoCrmGrademktMemberDetailDeleteAPIResponse.Get().(*TaobaoCrmGrademktMemberDetailDeleteAPIResponse)
}

// ReleaseTaobaoCrmGrademktMemberDetailDeleteAPIResponse 将 TaobaoCrmGrademktMemberDetailDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCrmGrademktMemberDetailDeleteAPIResponse(v *TaobaoCrmGrademktMemberDetailDeleteAPIResponse) {
	v.Reset()
	poolTaobaoCrmGrademktMemberDetailDeleteAPIResponse.Put(v)
}
