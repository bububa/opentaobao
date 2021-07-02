package foodscan

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFootscanMiniImageUploadAPIResponse 商家端图片上传 API返回值
// alibaba.footscan.mini.image.upload
//
// 提供图片上传功能，同时进行图片的检测
type AlibabaFootscanMiniImageUploadAPIResponse struct {
	model.CommonResponse
	AlibabaFootscanMiniImageUploadAPIResponseModel
}

// AlibabaFootscanMiniImageUploadAPIResponseModel is 商家端图片上传 成功返回结果
type AlibabaFootscanMiniImageUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_footscan_mini_image_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务出参
	Result *AlibabaFootscanMiniImageUploadMtopResult `json:"result,omitempty" xml:"result,omitempty"`
}
