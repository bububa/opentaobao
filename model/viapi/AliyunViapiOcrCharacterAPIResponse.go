package viapi

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AliyunViapiOcrCharacterAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliyunViapiOcrCharacterAPIResponseModel).Reset()
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

// Reset 清空结构体
func (m *AliyunViapiOcrCharacterAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TaobaoRequestId = ""
	m.Data = nil
}

var poolAliyunViapiOcrCharacterAPIResponse = sync.Pool{
	New: func() any {
		return new(AliyunViapiOcrCharacterAPIResponse)
	},
}

// GetAliyunViapiOcrCharacterAPIResponse 从 sync.Pool 获取 AliyunViapiOcrCharacterAPIResponse
func GetAliyunViapiOcrCharacterAPIResponse() *AliyunViapiOcrCharacterAPIResponse {
	return poolAliyunViapiOcrCharacterAPIResponse.Get().(*AliyunViapiOcrCharacterAPIResponse)
}

// ReleaseAliyunViapiOcrCharacterAPIResponse 将 AliyunViapiOcrCharacterAPIResponse 保存到 sync.Pool
func ReleaseAliyunViapiOcrCharacterAPIResponse(v *AliyunViapiOcrCharacterAPIResponse) {
	v.Reset()
	poolAliyunViapiOcrCharacterAPIResponse.Put(v)
}
