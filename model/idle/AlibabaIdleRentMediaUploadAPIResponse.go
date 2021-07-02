package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRentMediaUploadAPIResponse 闲鱼多媒体上传接口 API返回值
// alibaba.idle.rent.media.upload
//
// 上传多媒体信息，包括图片、视频（暂不支持）
type AlibabaIdleRentMediaUploadAPIResponse struct {
	model.CommonResponse
	AlibabaIdleRentMediaUploadAPIResponseModel
}

// AlibabaIdleRentMediaUploadAPIResponseModel is 闲鱼多媒体上传接口 成功返回结果
type AlibabaIdleRentMediaUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_rent_media_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应数据
	Result *AlibabaIdleRentMediaUploadTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
