package charity

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCharityCharitytimeUserQueryusercharityaccountAPIResponse 查询用户公益账户 API返回值
// alibaba.charity.charitytime.user.queryusercharityaccount
//
// 查询用户公益账户
type AlibabaCharityCharitytimeUserQueryusercharityaccountAPIResponse struct {
	model.CommonResponse
	AlibabaCharityCharitytimeUserQueryusercharityaccountAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCharityCharitytimeUserQueryusercharityaccountAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCharityCharitytimeUserQueryusercharityaccountAPIResponseModel).Reset()
}

// AlibabaCharityCharitytimeUserQueryusercharityaccountAPIResponseModel is 查询用户公益账户 成功返回结果
type AlibabaCharityCharitytimeUserQueryusercharityaccountAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_charity_charitytime_user_queryusercharityaccount_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应对象
	Result *CsrResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCharityCharitytimeUserQueryusercharityaccountAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCharityCharitytimeUserQueryusercharityaccountAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCharityCharitytimeUserQueryusercharityaccountAPIResponse)
	},
}

// GetAlibabaCharityCharitytimeUserQueryusercharityaccountAPIResponse 从 sync.Pool 获取 AlibabaCharityCharitytimeUserQueryusercharityaccountAPIResponse
func GetAlibabaCharityCharitytimeUserQueryusercharityaccountAPIResponse() *AlibabaCharityCharitytimeUserQueryusercharityaccountAPIResponse {
	return poolAlibabaCharityCharitytimeUserQueryusercharityaccountAPIResponse.Get().(*AlibabaCharityCharitytimeUserQueryusercharityaccountAPIResponse)
}

// ReleaseAlibabaCharityCharitytimeUserQueryusercharityaccountAPIResponse 将 AlibabaCharityCharitytimeUserQueryusercharityaccountAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCharityCharitytimeUserQueryusercharityaccountAPIResponse(v *AlibabaCharityCharitytimeUserQueryusercharityaccountAPIResponse) {
	v.Reset()
	poolAlibabaCharityCharitytimeUserQueryusercharityaccountAPIResponse.Put(v)
}
