package viapi

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliyunViapiOcrCharacterAPIResponse 通用文字识别 API返回值
// aliyun.viapi.ocr.character
//
// 获取通用的文字信息。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
type AliyunViapiOcrCharacterAPIResponse struct {
	model.CommonResponse
	AliyunViapiOcrCharacterAPIResponseModel
}

// AliyunViapiOcrCharacterAPIResponseModel is 通用文字识别 成功返回结果
type AliyunViapiOcrCharacterAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_viapi_ocr_character_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求ID
	TaobaoRequestId string `json:"taobao_request_id,omitempty" xml:"taobao_request_id,omitempty"`
	// 系统自动生成
	Data *AliyunViapiOcrCharacterData `json:"data,omitempty" xml:"data,omitempty"`
}
