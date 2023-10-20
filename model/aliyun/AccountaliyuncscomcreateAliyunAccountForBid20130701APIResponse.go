package aliyun

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AccountAliyuncsComCreateAliyunAccountForBid20130701APIResponse 为bid用户创建账号 API返回值
// account.aliyuncs.com.CreateAliyunAccountForBid.2013-07-01
//
// 给指定的bid创建账号，同时账号打上owner bid的标记
type AccountAliyuncsComCreateAliyunAccountForBid20130701APIResponse struct {
	model.CommonResponse
	AccountAliyuncsComCreateAliyunAccountForBid20130701APIResponseModel
}

// Reset 清空结构体
func (m *AccountAliyuncsComCreateAliyunAccountForBid20130701APIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AccountAliyuncsComCreateAliyunAccountForBid20130701APIResponseModel).Reset()
}

// AccountAliyuncsComCreateAliyunAccountForBid20130701APIResponseModel is 为bid用户创建账号 成功返回结果
type AccountAliyuncsComCreateAliyunAccountForBid20130701APIResponseModel struct {
	XMLName xml.Name `xml:"account_aliyuncs_com_CreateAliyunAccountForBid_2013-07-01_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ResultData string `json:"result_data,omitempty" xml:"result_data,omitempty"`
}

// Reset 清空结构体
func (m *AccountAliyuncsComCreateAliyunAccountForBid20130701APIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultData = ""
}

var poolAccountAliyuncsComCreateAliyunAccountForBid20130701APIResponse = sync.Pool{
	New: func() any {
		return new(AccountAliyuncsComCreateAliyunAccountForBid20130701APIResponse)
	},
}

// GetAccountAliyuncsComCreateAliyunAccountForBid20130701APIResponse 从 sync.Pool 获取 AccountAliyuncsComCreateAliyunAccountForBid20130701APIResponse
func GetAccountAliyuncsComCreateAliyunAccountForBid20130701APIResponse() *AccountAliyuncsComCreateAliyunAccountForBid20130701APIResponse {
	return poolAccountAliyuncsComCreateAliyunAccountForBid20130701APIResponse.Get().(*AccountAliyuncsComCreateAliyunAccountForBid20130701APIResponse)
}

// ReleaseAccountAliyuncsComCreateAliyunAccountForBid20130701APIResponse 将 AccountAliyuncsComCreateAliyunAccountForBid20130701APIResponse 保存到 sync.Pool
func ReleaseAccountAliyuncsComCreateAliyunAccountForBid20130701APIResponse(v *AccountAliyuncsComCreateAliyunAccountForBid20130701APIResponse) {
	v.Reset()
	poolAccountAliyuncsComCreateAliyunAccountForBid20130701APIResponse.Put(v)
}
