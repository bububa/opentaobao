package media

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPictureReplaceAPIResponse
替换图片 API返回值
taobao.picture.replace

替换一张图片，只替换图片数据，图片名称，图片分类等保持不变。 */
type TaobaoPictureReplaceAPIResponse struct {
	model.CommonResponse
	TaobaoPictureReplaceAPIResponseModel
}

// TaobaoPictureReplaceAPIResponseModel is 替换图片 成功返回结果
type TaobaoPictureReplaceAPIResponseModel struct {
	XMLName xml.Name `xml:"picture_replace_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 图片替换是否成功
	Done bool `json:"done,omitempty" xml:"done,omitempty"`
}
