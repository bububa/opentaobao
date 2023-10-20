package nlife

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNlifeB2cMemberDiscountruleGetAPIResponse 会员抵扣规则 API返回值
// alibaba.nlife.b2c.member.discountrule.get
//
// 获取企业会员抵扣规则
type AlibabaNlifeB2cMemberDiscountruleGetAPIResponse struct {
	model.CommonResponse
	AlibabaNlifeB2cMemberDiscountruleGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaNlifeB2cMemberDiscountruleGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaNlifeB2cMemberDiscountruleGetAPIResponseModel).Reset()
}

// AlibabaNlifeB2cMemberDiscountruleGetAPIResponseModel is 会员抵扣规则 成功返回结果
type AlibabaNlifeB2cMemberDiscountruleGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_nlife_b2c_member_discountrule_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结构化的文案
	DiscountMemos []DiscountMemo `json:"discount_memos,omitempty" xml:"discount_memos>discount_memo,omitempty"`
	// 错误码，当result为false时设置
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息，当result为false时设置
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 抵扣规则
	DiscountRule *DiscountRule `json:"discount_rule,omitempty" xml:"discount_rule,omitempty"`
	// 业务成功与否 true/false
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaNlifeB2cMemberDiscountruleGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DiscountMemos = m.DiscountMemos[:0]
	m.ErrCode = ""
	m.ErrMsg = ""
	m.DiscountRule = nil
	m.Result = false
}

var poolAlibabaNlifeB2cMemberDiscountruleGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaNlifeB2cMemberDiscountruleGetAPIResponse)
	},
}

// GetAlibabaNlifeB2cMemberDiscountruleGetAPIResponse 从 sync.Pool 获取 AlibabaNlifeB2cMemberDiscountruleGetAPIResponse
func GetAlibabaNlifeB2cMemberDiscountruleGetAPIResponse() *AlibabaNlifeB2cMemberDiscountruleGetAPIResponse {
	return poolAlibabaNlifeB2cMemberDiscountruleGetAPIResponse.Get().(*AlibabaNlifeB2cMemberDiscountruleGetAPIResponse)
}

// ReleaseAlibabaNlifeB2cMemberDiscountruleGetAPIResponse 将 AlibabaNlifeB2cMemberDiscountruleGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaNlifeB2cMemberDiscountruleGetAPIResponse(v *AlibabaNlifeB2cMemberDiscountruleGetAPIResponse) {
	v.Reset()
	poolAlibabaNlifeB2cMemberDiscountruleGetAPIResponse.Put(v)
}
