package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TmallProductSpecPicUploadAPIResponse
上传产品规格认证图片 API返回值
tmall.product.spec.pic.upload

上传指定类型的产品规格认证文件，并返回存有上传成功图片url的产品规格对象 */
type TmallProductSpecPicUploadAPIResponse struct {
	model.CommonResponse
	TmallProductSpecPicUploadAPIResponseModel
}

// TmallProductSpecPicUploadAPIResponseModel is 上传产品规格认证图片 成功返回结果
type TmallProductSpecPicUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_product_spec_pic_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 上传成功的产品规格认证图片url
	SpecPicUrl string `json:"spec_pic_url,omitempty" xml:"spec_pic_url,omitempty"`
}
