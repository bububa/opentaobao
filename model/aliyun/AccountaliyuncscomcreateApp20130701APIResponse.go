package aliyun

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AccountAliyuncsComCreateApp20130701APIResponse 给指定用户创建appkey API返回值
// account.aliyuncs.com.CreateApp.2013-07-01
//
// 为某个用户创建appkey
type AccountAliyuncsComCreateApp20130701APIResponse struct {
	model.CommonResponse
	AccountAliyuncsComCreateApp20130701APIResponseModel
}

// Reset 清空结构体
func (m *AccountAliyuncsComCreateApp20130701APIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AccountAliyuncsComCreateApp20130701APIResponseModel).Reset()
}

// AccountAliyuncsComCreateApp20130701APIResponseModel is 给指定用户创建appkey 成功返回结果
type AccountAliyuncsComCreateApp20130701APIResponseModel struct {
	XMLName xml.Name `xml:"account_aliyuncs_com_CreateApp_2013-07-01_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回调用信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	ResultData string `json:"result_data,omitempty" xml:"result_data,omitempty"`
}

// Reset 清空结构体
func (m *AccountAliyuncsComCreateApp20130701APIResponseModel) Reset() {
	m.RequestId = ""
	m.Code = ""
	m.Message = ""
	m.ResultData = ""
}

var poolAccountAliyuncsComCreateApp20130701APIResponse = sync.Pool{
	New: func() any {
		return new(AccountAliyuncsComCreateApp20130701APIResponse)
	},
}

// GetAccountAliyuncsComCreateApp20130701APIResponse 从 sync.Pool 获取 AccountAliyuncsComCreateApp20130701APIResponse
func GetAccountAliyuncsComCreateApp20130701APIResponse() *AccountAliyuncsComCreateApp20130701APIResponse {
	return poolAccountAliyuncsComCreateApp20130701APIResponse.Get().(*AccountAliyuncsComCreateApp20130701APIResponse)
}

// ReleaseAccountAliyuncsComCreateApp20130701APIResponse 将 AccountAliyuncsComCreateApp20130701APIResponse 保存到 sync.Pool
func ReleaseAccountAliyuncsComCreateApp20130701APIResponse(v *AccountAliyuncsComCreateApp20130701APIResponse) {
	v.Reset()
	poolAccountAliyuncsComCreateApp20130701APIResponse.Put(v)
}
