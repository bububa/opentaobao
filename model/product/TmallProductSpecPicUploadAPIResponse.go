package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallProductSpecPicUploadAPIResponse 上传产品规格认证图片 API返回值
// tmall.product.spec.pic.upload
//
// 上传指定类型的产品规格认证文件，并返回存有上传成功图片url的产品规格对象
type TmallProductSpecPicUploadAPIResponse struct {
	model.CommonResponse
	TmallProductSpecPicUploadAPIResponseModel
}

// Reset 清空结构体
func (m *TmallProductSpecPicUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallProductSpecPicUploadAPIResponseModel).Reset()
}

// TmallProductSpecPicUploadAPIResponseModel is 上传产品规格认证图片 成功返回结果
type TmallProductSpecPicUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_product_spec_pic_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 上传成功的产品规格认证图片url
	SpecPicUrl string `json:"spec_pic_url,omitempty" xml:"spec_pic_url,omitempty"`
}

// Reset 清空结构体
func (m *TmallProductSpecPicUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.SpecPicUrl = ""
}

var poolTmallProductSpecPicUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallProductSpecPicUploadAPIResponse)
	},
}

// GetTmallProductSpecPicUploadAPIResponse 从 sync.Pool 获取 TmallProductSpecPicUploadAPIResponse
func GetTmallProductSpecPicUploadAPIResponse() *TmallProductSpecPicUploadAPIResponse {
	return poolTmallProductSpecPicUploadAPIResponse.Get().(*TmallProductSpecPicUploadAPIResponse)
}

// ReleaseTmallProductSpecPicUploadAPIResponse 将 TmallProductSpecPicUploadAPIResponse 保存到 sync.Pool
func ReleaseTmallProductSpecPicUploadAPIResponse(v *TmallProductSpecPicUploadAPIResponse) {
	v.Reset()
	poolTmallProductSpecPicUploadAPIResponse.Put(v)
}
