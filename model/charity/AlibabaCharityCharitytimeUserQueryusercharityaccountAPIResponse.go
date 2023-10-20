package charity

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacharitycharitytimeuserqueryusercharityaccountAPIResponse 查询用户公益账户 API返回值
// alibaba.charity.charitytime.user.queryusercharityaccount
//
// 查询用户公益账户
type AlibabacharitycharitytimeuserqueryusercharityaccountAPIResponse struct {
	model.CommonResponse
	AlibabacharitycharitytimeuserqueryusercharityaccountAPIResponseModel
}

// AlibabacharitycharitytimeuserqueryusercharityaccountAPIResponseModel is 查询用户公益账户 成功返回结果
type AlibabacharitycharitytimeuserqueryusercharityaccountAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_charity_charitytime_user_queryusercharityaccount_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应对象
	Result *CsrResult `json:"result,omitempty" xml:"result,omitempty"`
}
