package xhotelcrm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelPotentialMemberBindAPIResponse 飞猪酒店商家会员绑定 API返回值
// taobao.xhotel.potential.member.bind
//
// 支持互通商家发起会员绑定
type TaobaoXhotelPotentialMemberBindAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelPotentialMemberBindAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelPotentialMemberBindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelPotentialMemberBindAPIResponseModel).Reset()
}

// TaobaoXhotelPotentialMemberBindAPIResponseModel is 飞猪酒店商家会员绑定 成功返回结果
type TaobaoXhotelPotentialMemberBindAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_potential_member_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 添加操作是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelPotentialMemberBindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoXhotelPotentialMemberBindAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelPotentialMemberBindAPIResponse)
	},
}

// GetTaobaoXhotelPotentialMemberBindAPIResponse 从 sync.Pool 获取 TaobaoXhotelPotentialMemberBindAPIResponse
func GetTaobaoXhotelPotentialMemberBindAPIResponse() *TaobaoXhotelPotentialMemberBindAPIResponse {
	return poolTaobaoXhotelPotentialMemberBindAPIResponse.Get().(*TaobaoXhotelPotentialMemberBindAPIResponse)
}

// ReleaseTaobaoXhotelPotentialMemberBindAPIResponse 将 TaobaoXhotelPotentialMemberBindAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelPotentialMemberBindAPIResponse(v *TaobaoXhotelPotentialMemberBindAPIResponse) {
	v.Reset()
	poolTaobaoXhotelPotentialMemberBindAPIResponse.Put(v)
}
