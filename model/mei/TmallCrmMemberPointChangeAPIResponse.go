package mei

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCrmMemberPointChangeAPIResponse 会员积分变更 API返回值
// tmall.crm.member.point.change
//
// 品牌CRM项目中，会员积分变更接口。
type TmallCrmMemberPointChangeAPIResponse struct {
	model.CommonResponse
	TmallCrmMemberPointChangeAPIResponseModel
}

// Reset 清空结构体
func (m *TmallCrmMemberPointChangeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallCrmMemberPointChangeAPIResponseModel).Reset()
}

// TmallCrmMemberPointChangeAPIResponseModel is 会员积分变更 成功返回结果
type TmallCrmMemberPointChangeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_crm_member_point_change_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用是否成功
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}

// Reset 清空结构体
func (m *TmallCrmMemberPointChangeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultSuccess = false
}

var poolTmallCrmMemberPointChangeAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallCrmMemberPointChangeAPIResponse)
	},
}

// GetTmallCrmMemberPointChangeAPIResponse 从 sync.Pool 获取 TmallCrmMemberPointChangeAPIResponse
func GetTmallCrmMemberPointChangeAPIResponse() *TmallCrmMemberPointChangeAPIResponse {
	return poolTmallCrmMemberPointChangeAPIResponse.Get().(*TmallCrmMemberPointChangeAPIResponse)
}

// ReleaseTmallCrmMemberPointChangeAPIResponse 将 TmallCrmMemberPointChangeAPIResponse 保存到 sync.Pool
func ReleaseTmallCrmMemberPointChangeAPIResponse(v *TmallCrmMemberPointChangeAPIResponse) {
	v.Reset()
	poolTmallCrmMemberPointChangeAPIResponse.Put(v)
}
