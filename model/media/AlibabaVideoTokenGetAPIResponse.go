package media

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabavideotokengetAPIResponse 获取上传token API返回值
// alibaba.video.token.get
//
// 获取上传token
type AlibabavideotokengetAPIResponse struct {
	model.CommonResponse
	AlibabavideotokengetAPIResponseModel
}

// AlibabavideotokengetAPIResponseModel is 获取上传token 成功返回结果
type AlibabavideotokengetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_video_token_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 上传token
	UploadToken string `json:"upload_token,omitempty" xml:"upload_token,omitempty"`
}
