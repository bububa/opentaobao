package mei

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallMeiCrmMemberGetbypaycodeAPIResponse 支付码获取会员信息 API返回值
// tmall.mei.crm.member.getbypaycode
//
// 通过支付码获取会员信息
type TmallMeiCrmMemberGetbypaycodeAPIResponse struct {
	model.CommonResponse
	TmallMeiCrmMemberGetbypaycodeAPIResponseModel
}

// Reset 清空结构体
func (m *TmallMeiCrmMemberGetbypaycodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallMeiCrmMemberGetbypaycodeAPIResponseModel).Reset()
}

// TmallMeiCrmMemberGetbypaycodeAPIResponseModel is 支付码获取会员信息 成功返回结果
type TmallMeiCrmMemberGetbypaycodeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_mei_crm_member_getbypaycode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallMeiCrmMemberGetbypaycodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallMeiCrmMemberGetbypaycodeAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallMeiCrmMemberGetbypaycodeAPIResponse)
	},
}

// GetTmallMeiCrmMemberGetbypaycodeAPIResponse 从 sync.Pool 获取 TmallMeiCrmMemberGetbypaycodeAPIResponse
func GetTmallMeiCrmMemberGetbypaycodeAPIResponse() *TmallMeiCrmMemberGetbypaycodeAPIResponse {
	return poolTmallMeiCrmMemberGetbypaycodeAPIResponse.Get().(*TmallMeiCrmMemberGetbypaycodeAPIResponse)
}

// ReleaseTmallMeiCrmMemberGetbypaycodeAPIResponse 将 TmallMeiCrmMemberGetbypaycodeAPIResponse 保存到 sync.Pool
func ReleaseTmallMeiCrmMemberGetbypaycodeAPIResponse(v *TmallMeiCrmMemberGetbypaycodeAPIResponse) {
	v.Reset()
	poolTmallMeiCrmMemberGetbypaycodeAPIResponse.Put(v)
}
