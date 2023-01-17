package util

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopSecretRegisterAPIResponse 注册加密账号 API返回值
// taobao.top.secret.register
//
// 提供给isv注册非淘系账号秘钥，isv依赖sdk自主加、解密
type TaobaoTopSecretRegisterAPIResponse struct {
	model.CommonResponse
	TaobaoTopSecretRegisterAPIResponseModel
}

// TaobaoTopSecretRegisterAPIResponseModel is 注册加密账号 成功返回结果
type TaobaoTopSecretRegisterAPIResponseModel struct {
	XMLName xml.Name `xml:"top_secret_register_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}
