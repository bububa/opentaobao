package user

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AccountAliyuncsComGetPubKey20130701APIResponse 获取用户公钥 API返回值
// account.aliyuncs.com.GetPubKey.2013-07-01
//
// 根据用户的appkey查询用户的pubkey
type AccountAliyuncsComGetPubKey20130701APIResponse struct {
	model.CommonResponse
	AccountAliyuncsComGetPubKey20130701APIResponseModel
}

// AccountAliyuncsComGetPubKey20130701APIResponseModel is 获取用户公钥 成功返回结果
type AccountAliyuncsComGetPubKey20130701APIResponseModel struct {
	XMLName xml.Name `xml:"account_aliyuncs_com_GetPubKey_2013-07-01_response"`
	// 用户的公钥
	PubKey string `json:"PubKey,omitempty" xml:"PubKey,omitempty"`
	// appkey的OwnerId
	OwnerId string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 用户appkey的类型
	KeyType string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	// 创建key的时间
	CreateTime string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// RequestId
	RequestId string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Message
	Message string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 返回code
	Code string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 用户的appkey
	AppKey string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}
