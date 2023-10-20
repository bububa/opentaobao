package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscUserCenterInfoQueryAPIResponse 查询授权的用户信息 API返回值
// alibaba.alsc.user.center.info.query
//
// 获取授权的饿了么用户信息
type AlibabaAlscUserCenterInfoQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlscUserCenterInfoQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscUserCenterInfoQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscUserCenterInfoQueryAPIResponseModel).Reset()
}

// AlibabaAlscUserCenterInfoQueryAPIResponseModel is 查询授权的用户信息 成功返回结果
type AlibabaAlscUserCenterInfoQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_user_center_info_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 脱敏昵称
	EncryptNickName string `json:"encrypt_nick_name,omitempty" xml:"encrypt_nick_name,omitempty"`
	// 昵称
	NickName string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	// 简介
	Intro string `json:"intro,omitempty" xml:"intro,omitempty"`
	// 手机号
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 头像链接
	AccessAvatarUrl string `json:"access_avatar_url,omitempty" xml:"access_avatar_url,omitempty"`
	// 用户名
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// 脱敏手机号
	EncryptMobile string `json:"encrypt_mobile,omitempty" xml:"encrypt_mobile,omitempty"`
	// 脱敏用户名
	EncryptUsername string `json:"encrypt_username,omitempty" xml:"encrypt_username,omitempty"`
	// 不同appkey下同一用户的openId不同
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscUserCenterInfoQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.EncryptNickName = ""
	m.NickName = ""
	m.Intro = ""
	m.Mobile = ""
	m.AccessAvatarUrl = ""
	m.UserName = ""
	m.EncryptMobile = ""
	m.EncryptUsername = ""
	m.OpenId = ""
}

var poolAlibabaAlscUserCenterInfoQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscUserCenterInfoQueryAPIResponse)
	},
}

// GetAlibabaAlscUserCenterInfoQueryAPIResponse 从 sync.Pool 获取 AlibabaAlscUserCenterInfoQueryAPIResponse
func GetAlibabaAlscUserCenterInfoQueryAPIResponse() *AlibabaAlscUserCenterInfoQueryAPIResponse {
	return poolAlibabaAlscUserCenterInfoQueryAPIResponse.Get().(*AlibabaAlscUserCenterInfoQueryAPIResponse)
}

// ReleaseAlibabaAlscUserCenterInfoQueryAPIResponse 将 AlibabaAlscUserCenterInfoQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscUserCenterInfoQueryAPIResponse(v *AlibabaAlscUserCenterInfoQueryAPIResponse) {
	v.Reset()
	poolAlibabaAlscUserCenterInfoQueryAPIResponse.Put(v)
}
