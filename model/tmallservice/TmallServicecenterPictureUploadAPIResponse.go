package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterpictureuploadAPIResponse 上传图片 API返回值
// tmall.servicecenter.picture.upload
//
// 给服务商ERP系统使用，用于上传图片保存在天猫，一般用于工单信息回传时候保存服务商的服务证明信息相关的图片。
type TmallservicecenterpictureuploadAPIResponse struct {
	model.CommonResponse
	TmallservicecenterpictureuploadAPIResponseModel
}

// TmallservicecenterpictureuploadAPIResponseModel is 上传图片 成功返回结果
type TmallservicecenterpictureuploadAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_picture_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}
