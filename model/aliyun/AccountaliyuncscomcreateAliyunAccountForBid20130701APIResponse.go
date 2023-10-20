package aliyun

import (
	"encoding/xml"

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

// AccountAliyuncsComCreateAliyunAccountForBid20130701APIResponseModel is 为bid用户创建账号 成功返回结果
type AccountAliyuncsComCreateAliyunAccountForBid20130701APIResponseModel struct {
	XMLName xml.Name `xml:"account_aliyuncs_com_CreateAliyunAccountForBid_2013-07-01_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ResultData string `json:"result_data,omitempty" xml:"result_data,omitempty"`
}
