package aliyun

import (
	"encoding/xml"

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

// AccountAliyuncsComCreateAppForBid20130701APIResponseModel is 给当前渠道下的用户创建app 成功返回结果
type AccountAliyuncsComCreateAppForBid20130701APIResponseModel struct {
	XMLName xml.Name `xml:"account_aliyuncs_com_CreateAppForBid_2013-07-01_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ResultData string `json:"result_data,omitempty" xml:"result_data,omitempty"`
}
