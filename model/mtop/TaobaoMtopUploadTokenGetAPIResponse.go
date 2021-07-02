package mtop

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMtopUploadTokenGetAPIResponse 获取文件上传授权 API返回值
// taobao.mtop.upload.token.get
//
// 获取mtop文件上传授权
type TaobaoMtopUploadTokenGetAPIResponse struct {
	model.CommonResponse
	TaobaoMtopUploadTokenGetAPIResponseModel
}

// TaobaoMtopUploadTokenGetAPIResponseModel is 获取文件上传授权 成功返回结果
type TaobaoMtopUploadTokenGetAPIResponseModel struct {
	XMLName xml.Name `xml:"mtop_upload_token_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msg
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 单次上传文件块最大大小，单位 byte
	MaxBodyLength int64 `json:"max_body_length,omitempty" xml:"max_body_length,omitempty"`
	// 单个文件重试上传次数
	MaxRetryTimes int64 `json:"max_retry_times,omitempty" xml:"max_retry_times,omitempty"`
	// 本次指定的上传文件服务器地址
	ServerAddress string `json:"server_address,omitempty" xml:"server_address,omitempty"`
	// token失效时间点
	Timeout int64 `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// 颁发的上传令牌
	Token string `json:"token,omitempty" xml:"token,omitempty"`
}
