package mei

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCrmMemberFrontUnbindPrivyAPIResponse 品牌会员解绑(隐私号版） API返回值
// tmall.crm.member.front.unbind.privy
//
// 品牌会员解绑功能
type TmallCrmMemberFrontUnbindPrivyAPIResponse struct {
	model.CommonResponse
	TmallCrmMemberFrontUnbindPrivyAPIResponseModel
}

// Reset 清空结构体
func (m *TmallCrmMemberFrontUnbindPrivyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallCrmMemberFrontUnbindPrivyAPIResponseModel).Reset()
}

// TmallCrmMemberFrontUnbindPrivyAPIResponseModel is 品牌会员解绑(隐私号版） 成功返回结果
type TmallCrmMemberFrontUnbindPrivyAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_crm_member_front_unbind_privy_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口调用是否成功
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}

// Reset 清空结构体
func (m *TmallCrmMemberFrontUnbindPrivyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultSuccess = false
}

var poolTmallCrmMemberFrontUnbindPrivyAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallCrmMemberFrontUnbindPrivyAPIResponse)
	},
}

// GetTmallCrmMemberFrontUnbindPrivyAPIResponse 从 sync.Pool 获取 TmallCrmMemberFrontUnbindPrivyAPIResponse
func GetTmallCrmMemberFrontUnbindPrivyAPIResponse() *TmallCrmMemberFrontUnbindPrivyAPIResponse {
	return poolTmallCrmMemberFrontUnbindPrivyAPIResponse.Get().(*TmallCrmMemberFrontUnbindPrivyAPIResponse)
}

// ReleaseTmallCrmMemberFrontUnbindPrivyAPIResponse 将 TmallCrmMemberFrontUnbindPrivyAPIResponse 保存到 sync.Pool
func ReleaseTmallCrmMemberFrontUnbindPrivyAPIResponse(v *TmallCrmMemberFrontUnbindPrivyAPIResponse) {
	v.Reset()
	poolTmallCrmMemberFrontUnbindPrivyAPIResponse.Put(v)
}
