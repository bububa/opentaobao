package jym

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJymMemberRealnameVerifyWithoutuidAPIResponse 用户实名认证 API返回值
// taobao.jym.member.realname.verify.withoutuid
//
// 开放用户实名认证接口使用
type TaobaoJymMemberRealnameVerifyWithoutuidAPIResponse struct {
	model.CommonResponse
	TaobaoJymMemberRealnameVerifyWithoutuidAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJymMemberRealnameVerifyWithoutuidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJymMemberRealnameVerifyWithoutuidAPIResponseModel).Reset()
}

// TaobaoJymMemberRealnameVerifyWithoutuidAPIResponseModel is 用户实名认证 成功返回结果
type TaobaoJymMemberRealnameVerifyWithoutuidAPIResponseModel struct {
	XMLName xml.Name `xml:"jym_member_realname_verify_withoutuid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 实名认证结果
	Result *TaobaoJymMemberRealnameVerifyWithoutuidResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJymMemberRealnameVerifyWithoutuidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoJymMemberRealnameVerifyWithoutuidAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJymMemberRealnameVerifyWithoutuidAPIResponse)
	},
}

// GetTaobaoJymMemberRealnameVerifyWithoutuidAPIResponse 从 sync.Pool 获取 TaobaoJymMemberRealnameVerifyWithoutuidAPIResponse
func GetTaobaoJymMemberRealnameVerifyWithoutuidAPIResponse() *TaobaoJymMemberRealnameVerifyWithoutuidAPIResponse {
	return poolTaobaoJymMemberRealnameVerifyWithoutuidAPIResponse.Get().(*TaobaoJymMemberRealnameVerifyWithoutuidAPIResponse)
}

// ReleaseTaobaoJymMemberRealnameVerifyWithoutuidAPIResponse 将 TaobaoJymMemberRealnameVerifyWithoutuidAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJymMemberRealnameVerifyWithoutuidAPIResponse(v *TaobaoJymMemberRealnameVerifyWithoutuidAPIResponse) {
	v.Reset()
	poolTaobaoJymMemberRealnameVerifyWithoutuidAPIResponse.Put(v)
}
