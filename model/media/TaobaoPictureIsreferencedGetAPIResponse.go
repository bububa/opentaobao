package media

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaopictureisreferencedgetAPIResponse 图片是否被引用 API返回值
// taobao.picture.isreferenced.get
//
// 查询图片是否被引用，被引用返回true，未被引用返回false
type TaobaopictureisreferencedgetAPIResponse struct {
	model.CommonResponse
	TaobaopictureisreferencedgetAPIResponseModel
}

// TaobaopictureisreferencedgetAPIResponseModel is 图片是否被引用 成功返回结果
type TaobaopictureisreferencedgetAPIResponseModel struct {
	XMLName xml.Name `xml:"picture_isreferenced_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 图片是否被引用
	IsReferenced bool `json:"is_referenced,omitempty" xml:"is_referenced,omitempty"`
}
