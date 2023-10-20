package media

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaopictureuploadAPIResponse 上传单张图片 API返回值
// taobao.picture.upload
//
// 图片空间上传接口
type TaobaopictureuploadAPIResponse struct {
	model.CommonResponse
	TaobaopictureuploadAPIResponseModel
}

// TaobaopictureuploadAPIResponseModel is 上传单张图片 成功返回结果
type TaobaopictureuploadAPIResponseModel struct {
	XMLName xml.Name `xml:"picture_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 当前上传的一张图片信息
	Picture *Picture `json:"picture,omitempty" xml:"picture,omitempty"`
}
