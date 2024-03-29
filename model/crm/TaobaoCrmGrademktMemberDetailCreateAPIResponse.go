package crm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmGrademktMemberDetailCreateAPIResponse 会员等级营销-创建商品等级营销明细 API返回值
// taobao.crm.grademkt.member.detail.create
//
// 创建商品等级营销明细
type TaobaoCrmGrademktMemberDetailCreateAPIResponse struct {
	model.CommonResponse
	TaobaoCrmGrademktMemberDetailCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCrmGrademktMemberDetailCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCrmGrademktMemberDetailCreateAPIResponseModel).Reset()
}

// TaobaoCrmGrademktMemberDetailCreateAPIResponseModel is 会员等级营销-创建商品等级营销明细 成功返回结果
type TaobaoCrmGrademktMemberDetailCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_grademkt_member_detail_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// json格式
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCrmGrademktMemberDetailCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Module = false
}

var poolTaobaoCrmGrademktMemberDetailCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCrmGrademktMemberDetailCreateAPIResponse)
	},
}

// GetTaobaoCrmGrademktMemberDetailCreateAPIResponse 从 sync.Pool 获取 TaobaoCrmGrademktMemberDetailCreateAPIResponse
func GetTaobaoCrmGrademktMemberDetailCreateAPIResponse() *TaobaoCrmGrademktMemberDetailCreateAPIResponse {
	return poolTaobaoCrmGrademktMemberDetailCreateAPIResponse.Get().(*TaobaoCrmGrademktMemberDetailCreateAPIResponse)
}

// ReleaseTaobaoCrmGrademktMemberDetailCreateAPIResponse 将 TaobaoCrmGrademktMemberDetailCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCrmGrademktMemberDetailCreateAPIResponse(v *TaobaoCrmGrademktMemberDetailCreateAPIResponse) {
	v.Reset()
	poolTaobaoCrmGrademktMemberDetailCreateAPIResponse.Put(v)
}
