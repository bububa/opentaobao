package aliyun

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AccountAliyuncsComDeleteAppForBid20130701APIResponse 运营商删除用户的appkey API返回值
// account.aliyuncs.com.DeleteAppForBid.2013-07-01
//
// 删除用户的appkey，会校验调用的用户是否为当前运营商所创建的。
type AccountAliyuncsComDeleteAppForBid20130701APIResponse struct {
	model.CommonResponse
	AccountAliyuncsComDeleteAppForBid20130701APIResponseModel
}

// Reset 清空结构体
func (m *AccountAliyuncsComDeleteAppForBid20130701APIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AccountAliyuncsComDeleteAppForBid20130701APIResponseModel).Reset()
}

// AccountAliyuncsComDeleteAppForBid20130701APIResponseModel is 运营商删除用户的appkey 成功返回结果
type AccountAliyuncsComDeleteAppForBid20130701APIResponseModel struct {
	XMLName xml.Name `xml:"account_aliyuncs_com_DeleteAppForBid_2013-07-01_response"`
	// 用户删除的appkey
	AppKey string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// 调用描述信息
	Message string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 结果code
	Code string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 入参的requestId
	RequestId string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

// Reset 清空结构体
func (m *AccountAliyuncsComDeleteAppForBid20130701APIResponseModel) Reset() {
	m.RequestId = ""
	m.AppKey = ""
	m.Message = ""
	m.Code = ""
	m.RequestId = ""
}

var poolAccountAliyuncsComDeleteAppForBid20130701APIResponse = sync.Pool{
	New: func() any {
		return new(AccountAliyuncsComDeleteAppForBid20130701APIResponse)
	},
}

// GetAccountAliyuncsComDeleteAppForBid20130701APIResponse 从 sync.Pool 获取 AccountAliyuncsComDeleteAppForBid20130701APIResponse
func GetAccountAliyuncsComDeleteAppForBid20130701APIResponse() *AccountAliyuncsComDeleteAppForBid20130701APIResponse {
	return poolAccountAliyuncsComDeleteAppForBid20130701APIResponse.Get().(*AccountAliyuncsComDeleteAppForBid20130701APIResponse)
}

// ReleaseAccountAliyuncsComDeleteAppForBid20130701APIResponse 将 AccountAliyuncsComDeleteAppForBid20130701APIResponse 保存到 sync.Pool
func ReleaseAccountAliyuncsComDeleteAppForBid20130701APIResponse(v *AccountAliyuncsComDeleteAppForBid20130701APIResponse) {
	v.Reset()
	poolAccountAliyuncsComDeleteAppForBid20130701APIResponse.Put(v)
}
