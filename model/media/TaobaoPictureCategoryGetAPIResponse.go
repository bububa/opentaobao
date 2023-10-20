package media

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaopicturecategorygetAPIResponse 获取图片分类信息 API返回值
// taobao.picture.category.get
//
// 获取图片分类信息
type TaobaopicturecategorygetAPIResponse struct {
	model.CommonResponse
	TaobaopicturecategorygetAPIResponseModel
}

// TaobaopicturecategorygetAPIResponseModel is 获取图片分类信息 成功返回结果
type TaobaopicturecategorygetAPIResponseModel struct {
	XMLName xml.Name `xml:"picture_category_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 图片分类
	PictureCategories []PictureCategory `json:"picture_categories,omitempty" xml:"picture_categories>picture_category,omitempty"`
}
