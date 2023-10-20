package aliyun

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AccountAliyuncsComCreateAliyunAccount20130701APIResponse 创建阿里云账号 API返回值
// account.aliyuncs.com.CreateAliyunAccount.2013-07-01
//
// 根据给定的阿里云账号，密码以及手机号创建阿里云账号
type AccountAliyuncsComCreateAliyunAccount20130701APIResponse struct {
	model.CommonResponse
	AccountAliyuncsComCreateAliyunAccount20130701APIResponseModel
}

// Reset 清空结构体
func (m *AccountAliyuncsComCreateAliyunAccount20130701APIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AccountAliyuncsComCreateAliyunAccount20130701APIResponseModel).Reset()
}

// AccountAliyuncsComCreateAliyunAccount20130701APIResponseModel is 创建阿里云账号 成功返回结果
type AccountAliyuncsComCreateAliyunAccount20130701APIResponseModel struct {
	XMLName xml.Name `xml:"account_aliyuncs_com_CreateAliyunAccount_2013-07-01_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ResultData string `json:"result_data,omitempty" xml:"result_data,omitempty"`
}

// Reset 清空结构体
func (m *AccountAliyuncsComCreateAliyunAccount20130701APIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultData = ""
}

var poolAccountAliyuncsComCreateAliyunAccount20130701APIResponse = sync.Pool{
	New: func() any {
		return new(AccountAliyuncsComCreateAliyunAccount20130701APIResponse)
	},
}

// GetAccountAliyuncsComCreateAliyunAccount20130701APIResponse 从 sync.Pool 获取 AccountAliyuncsComCreateAliyunAccount20130701APIResponse
func GetAccountAliyuncsComCreateAliyunAccount20130701APIResponse() *AccountAliyuncsComCreateAliyunAccount20130701APIResponse {
	return poolAccountAliyuncsComCreateAliyunAccount20130701APIResponse.Get().(*AccountAliyuncsComCreateAliyunAccount20130701APIResponse)
}

// ReleaseAccountAliyuncsComCreateAliyunAccount20130701APIResponse 将 AccountAliyuncsComCreateAliyunAccount20130701APIResponse 保存到 sync.Pool
func ReleaseAccountAliyuncsComCreateAliyunAccount20130701APIResponse(v *AccountAliyuncsComCreateAliyunAccount20130701APIResponse) {
	v.Reset()
	poolAccountAliyuncsComCreateAliyunAccount20130701APIResponse.Put(v)
}
