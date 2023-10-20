package aliyun

import (
	"encoding/xml"

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

// AccountAliyuncsComCreateAliyunAccount20130701APIResponseModel is 创建阿里云账号 成功返回结果
type AccountAliyuncsComCreateAliyunAccount20130701APIResponseModel struct {
	XMLName xml.Name `xml:"account_aliyuncs_com_CreateAliyunAccount_2013-07-01_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ResultData string `json:"result_data,omitempty" xml:"result_data,omitempty"`
}
