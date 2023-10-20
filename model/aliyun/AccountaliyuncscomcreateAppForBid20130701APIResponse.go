package aliyun

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AccountAliyuncsComCreateAppForBid20130701APIResponse 给当前渠道下的用户创建app API返回值
// account.aliyuncs.com.CreateAppForBid.2013-07-01
//
// 给自己渠道下的用户创建app
type AccountAliyuncsComCreateAppForBid20130701APIResponse struct {
	model.CommonResponse
	AccountAliyuncsComCreateAppForBid20130701APIResponseModel
}

// Reset 清空结构体
func (m *AccountAliyuncsComCreateAppForBid20130701APIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AccountAliyuncsComCreateAppForBid20130701APIResponseModel).Reset()
}

// AccountAliyuncsComCreateAppForBid20130701APIResponseModel is 给当前渠道下的用户创建app 成功返回结果
type AccountAliyuncsComCreateAppForBid20130701APIResponseModel struct {
	XMLName xml.Name `xml:"account_aliyuncs_com_CreateAppForBid_2013-07-01_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ResultData string `json:"result_data,omitempty" xml:"result_data,omitempty"`
}

// Reset 清空结构体
func (m *AccountAliyuncsComCreateAppForBid20130701APIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultData = ""
}

var poolAccountAliyuncsComCreateAppForBid20130701APIResponse = sync.Pool{
	New: func() any {
		return new(AccountAliyuncsComCreateAppForBid20130701APIResponse)
	},
}

// GetAccountAliyuncsComCreateAppForBid20130701APIResponse 从 sync.Pool 获取 AccountAliyuncsComCreateAppForBid20130701APIResponse
func GetAccountAliyuncsComCreateAppForBid20130701APIResponse() *AccountAliyuncsComCreateAppForBid20130701APIResponse {
	return poolAccountAliyuncsComCreateAppForBid20130701APIResponse.Get().(*AccountAliyuncsComCreateAppForBid20130701APIResponse)
}

// ReleaseAccountAliyuncsComCreateAppForBid20130701APIResponse 将 AccountAliyuncsComCreateAppForBid20130701APIResponse 保存到 sync.Pool
func ReleaseAccountAliyuncsComCreateAppForBid20130701APIResponse(v *AccountAliyuncsComCreateAppForBid20130701APIResponse) {
	v.Reset()
	poolAccountAliyuncsComCreateAppForBid20130701APIResponse.Put(v)
}
