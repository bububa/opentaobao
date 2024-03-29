package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMemberCardGetAPIResponse 查询会员卡信息 API返回值
// alibaba.wdk.member.card.get
//
// 根据会员卡查询会员信息
type AlibabaWdkMemberCardGetAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMemberCardGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMemberCardGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMemberCardGetAPIResponseModel).Reset()
}

// AlibabaWdkMemberCardGetAPIResponseModel is 查询会员卡信息 成功返回结果
type AlibabaWdkMemberCardGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_member_card_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	ApiResult *AlibabaWdkMemberCardGetApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMemberCardGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaWdkMemberCardGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMemberCardGetAPIResponse)
	},
}

// GetAlibabaWdkMemberCardGetAPIResponse 从 sync.Pool 获取 AlibabaWdkMemberCardGetAPIResponse
func GetAlibabaWdkMemberCardGetAPIResponse() *AlibabaWdkMemberCardGetAPIResponse {
	return poolAlibabaWdkMemberCardGetAPIResponse.Get().(*AlibabaWdkMemberCardGetAPIResponse)
}

// ReleaseAlibabaWdkMemberCardGetAPIResponse 将 AlibabaWdkMemberCardGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMemberCardGetAPIResponse(v *AlibabaWdkMemberCardGetAPIResponse) {
	v.Reset()
	poolAlibabaWdkMemberCardGetAPIResponse.Put(v)
}
