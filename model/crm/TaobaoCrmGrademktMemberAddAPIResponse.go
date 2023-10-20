package crm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmGrademktMemberAddAPIResponse 会员等级营销-会员吸纳 API返回值
// taobao.crm.grademkt.member.add
//
// 商家通过该接口吸纳线上店铺会员。
type TaobaoCrmGrademktMemberAddAPIResponse struct {
	model.CommonResponse
	TaobaoCrmGrademktMemberAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCrmGrademktMemberAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCrmGrademktMemberAddAPIResponseModel).Reset()
}

// TaobaoCrmGrademktMemberAddAPIResponseModel is 会员等级营销-会员吸纳 成功返回结果
type TaobaoCrmGrademktMemberAddAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_grademkt_member_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回操作是否成功
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCrmGrademktMemberAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = false
}

var poolTaobaoCrmGrademktMemberAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCrmGrademktMemberAddAPIResponse)
	},
}

// GetTaobaoCrmGrademktMemberAddAPIResponse 从 sync.Pool 获取 TaobaoCrmGrademktMemberAddAPIResponse
func GetTaobaoCrmGrademktMemberAddAPIResponse() *TaobaoCrmGrademktMemberAddAPIResponse {
	return poolTaobaoCrmGrademktMemberAddAPIResponse.Get().(*TaobaoCrmGrademktMemberAddAPIResponse)
}

// ReleaseTaobaoCrmGrademktMemberAddAPIResponse 将 TaobaoCrmGrademktMemberAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCrmGrademktMemberAddAPIResponse(v *TaobaoCrmGrademktMemberAddAPIResponse) {
	v.Reset()
	poolTaobaoCrmGrademktMemberAddAPIResponse.Put(v)
}
