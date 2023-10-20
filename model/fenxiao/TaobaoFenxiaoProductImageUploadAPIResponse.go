package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaoproductimageuploadAPIResponse 产品图片上传 API返回值
// taobao.fenxiao.product.image.upload
//
// 产品主图图片空间相对路径或绝对路径添加或更新，或者是图片上传。如果指定位置的图片已存在，则覆盖原有信息。如果位置为1,自动设为主图；如果位置为0，表示属性图片
type TaobaofenxiaoproductimageuploadAPIResponse struct {
	model.CommonResponse
	TaobaofenxiaoproductimageuploadAPIResponseModel
}

// TaobaofenxiaoproductimageuploadAPIResponseModel is 产品图片上传 成功返回结果
type TaobaofenxiaoproductimageuploadAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_product_image_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作时间
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 操作是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
