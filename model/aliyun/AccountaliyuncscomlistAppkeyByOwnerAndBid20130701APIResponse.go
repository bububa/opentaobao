package aliyun

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AccountAliyuncsComListAppkeyByOwnerAndBid20130701APIResponse 根据渠道id和状态列出appkey API返回值
// account.aliyuncs.com.ListAppkeyByOwnerAndBid.2013-07-01
//
// 根据渠道id和状态列出appkey
type AccountAliyuncsComListAppkeyByOwnerAndBid20130701APIResponse struct {
	model.CommonResponse
	AccountAliyuncsComListAppkeyByOwnerAndBid20130701APIResponseModel
}

// Reset 清空结构体
func (m *AccountAliyuncsComListAppkeyByOwnerAndBid20130701APIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AccountAliyuncsComListAppkeyByOwnerAndBid20130701APIResponseModel).Reset()
}

// AccountAliyuncsComListAppkeyByOwnerAndBid20130701APIResponseModel is 根据渠道id和状态列出appkey 成功返回结果
type AccountAliyuncsComListAppkeyByOwnerAndBid20130701APIResponseModel struct {
	XMLName xml.Name `xml:"account_aliyuncs_com_ListAppkeyByOwnerAndBid_2013-07-01_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// return result
	ResultData string `json:"result_data,omitempty" xml:"result_data,omitempty"`
}

// Reset 清空结构体
func (m *AccountAliyuncsComListAppkeyByOwnerAndBid20130701APIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultData = ""
}

var poolAccountAliyuncsComListAppkeyByOwnerAndBid20130701APIResponse = sync.Pool{
	New: func() any {
		return new(AccountAliyuncsComListAppkeyByOwnerAndBid20130701APIResponse)
	},
}

// GetAccountAliyuncsComListAppkeyByOwnerAndBid20130701APIResponse 从 sync.Pool 获取 AccountAliyuncsComListAppkeyByOwnerAndBid20130701APIResponse
func GetAccountAliyuncsComListAppkeyByOwnerAndBid20130701APIResponse() *AccountAliyuncsComListAppkeyByOwnerAndBid20130701APIResponse {
	return poolAccountAliyuncsComListAppkeyByOwnerAndBid20130701APIResponse.Get().(*AccountAliyuncsComListAppkeyByOwnerAndBid20130701APIResponse)
}

// ReleaseAccountAliyuncsComListAppkeyByOwnerAndBid20130701APIResponse 将 AccountAliyuncsComListAppkeyByOwnerAndBid20130701APIResponse 保存到 sync.Pool
func ReleaseAccountAliyuncsComListAppkeyByOwnerAndBid20130701APIResponse(v *AccountAliyuncsComListAppkeyByOwnerAndBid20130701APIResponse) {
	v.Reset()
	poolAccountAliyuncsComListAppkeyByOwnerAndBid20130701APIResponse.Put(v)
}
