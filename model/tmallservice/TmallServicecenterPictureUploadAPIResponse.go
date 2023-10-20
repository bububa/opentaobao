package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterPictureUploadAPIResponse 上传图片 API返回值
// tmall.servicecenter.picture.upload
//
// 给服务商ERP系统使用，用于上传图片保存在天猫，一般用于工单信息回传时候保存服务商的服务证明信息相关的图片。
type TmallServicecenterPictureUploadAPIResponse struct {
	model.CommonResponse
	TmallServicecenterPictureUploadAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterPictureUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterPictureUploadAPIResponseModel).Reset()
}

// TmallServicecenterPictureUploadAPIResponseModel is 上传图片 成功返回结果
type TmallServicecenterPictureUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_picture_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterPictureUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterPictureUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterPictureUploadAPIResponse)
	},
}

// GetTmallServicecenterPictureUploadAPIResponse 从 sync.Pool 获取 TmallServicecenterPictureUploadAPIResponse
func GetTmallServicecenterPictureUploadAPIResponse() *TmallServicecenterPictureUploadAPIResponse {
	return poolTmallServicecenterPictureUploadAPIResponse.Get().(*TmallServicecenterPictureUploadAPIResponse)
}

// ReleaseTmallServicecenterPictureUploadAPIResponse 将 TmallServicecenterPictureUploadAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterPictureUploadAPIResponse(v *TmallServicecenterPictureUploadAPIResponse) {
	v.Reset()
	poolTmallServicecenterPictureUploadAPIResponse.Put(v)
}
