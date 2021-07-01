package ihome

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIhomeCtomContentImgUploadAPIResponse
实拍图投稿链路上传图片 API返回值
alibaba.ihome.ctom.content.img.upload

实拍图投稿链路上传图片 */
type AlibabaIhomeCtomContentImgUploadAPIResponse struct {
	model.CommonResponse
	AlibabaIhomeCtomContentImgUploadAPIResponseModel
}

// AlibabaIhomeCtomContentImgUploadAPIResponseModel is 实拍图投稿链路上传图片 成功返回结果
type AlibabaIhomeCtomContentImgUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ihome_ctom_content_img_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 根据站点名称查询产品
	Result *AlibabaIhomeCtomContentImgUploadApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
