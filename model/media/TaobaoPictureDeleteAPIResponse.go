package media

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPictureDeleteAPIResponse 删除图片空间图片 API返回值
// taobao.picture.delete
//
// 删除图片空间图片
type TaobaoPictureDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoPictureDeleteAPIResponseModel
}

// TaobaoPictureDeleteAPIResponseModel is 删除图片空间图片 成功返回结果
type TaobaoPictureDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"picture_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否删除
	Success string `json:"success,omitempty" xml:"success,omitempty"`
}
