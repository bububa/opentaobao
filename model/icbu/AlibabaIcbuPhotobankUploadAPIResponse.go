package icbu

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuPhotobankUploadAPIResponse 图片银行图片上传开放接口 API返回值
// alibaba.icbu.photobank.upload
//
// 图片银行图片上传开放接口
type AlibabaIcbuPhotobankUploadAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuPhotobankUploadAPIResponseModel
}

// AlibabaIcbuPhotobankUploadAPIResponseModel is 图片银行图片上传开放接口 成功返回结果
type AlibabaIcbuPhotobankUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_photobank_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 图片信息
	UploadImageResponse *UploadImageResponseDo `json:"upload_image_response,omitempty" xml:"upload_image_response,omitempty"`
}
