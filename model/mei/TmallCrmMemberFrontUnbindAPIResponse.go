package mei

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCrmMemberFrontUnbindAPIResponse 品牌会员解绑 API返回值
// tmall.crm.member.front.unbind
//
// 品牌会员解绑功能
type TmallCrmMemberFrontUnbindAPIResponse struct {
	model.CommonResponse
	TmallCrmMemberFrontUnbindAPIResponseModel
}

// Reset 清空结构体
func (m *TmallCrmMemberFrontUnbindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallCrmMemberFrontUnbindAPIResponseModel).Reset()
}

// TmallCrmMemberFrontUnbindAPIResponseModel is 品牌会员解绑 成功返回结果
type TmallCrmMemberFrontUnbindAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_crm_member_front_unbind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口调用是否成功
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}

// Reset 清空结构体
func (m *TmallCrmMemberFrontUnbindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultSuccess = false
}

var poolTmallCrmMemberFrontUnbindAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallCrmMemberFrontUnbindAPIResponse)
	},
}

// GetTmallCrmMemberFrontUnbindAPIResponse 从 sync.Pool 获取 TmallCrmMemberFrontUnbindAPIResponse
func GetTmallCrmMemberFrontUnbindAPIResponse() *TmallCrmMemberFrontUnbindAPIResponse {
	return poolTmallCrmMemberFrontUnbindAPIResponse.Get().(*TmallCrmMemberFrontUnbindAPIResponse)
}

// ReleaseTmallCrmMemberFrontUnbindAPIResponse 将 TmallCrmMemberFrontUnbindAPIResponse 保存到 sync.Pool
func ReleaseTmallCrmMemberFrontUnbindAPIResponse(v *TmallCrmMemberFrontUnbindAPIResponse) {
	v.Reset()
	poolTmallCrmMemberFrontUnbindAPIResponse.Put(v)
}
