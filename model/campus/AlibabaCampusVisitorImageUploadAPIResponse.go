package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusVisitorImageUploadAPIResponse 访客大厅图片上传及查看 API返回值
// alibaba.campus.visitor.image.upload
//
// 访客大厅图片上传及查看
type AlibabaCampusVisitorImageUploadAPIResponse struct {
	model.CommonResponse
	AlibabaCampusVisitorImageUploadAPIResponseModel
}

// AlibabaCampusVisitorImageUploadAPIResponseModel is 访客大厅图片上传及查看 成功返回结果
type AlibabaCampusVisitorImageUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_visitor_image_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
