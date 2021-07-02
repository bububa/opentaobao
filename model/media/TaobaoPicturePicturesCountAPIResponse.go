package media

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPicturePicturesCountAPIResponse 图片总数查询 API返回值
// taobao.picture.pictures.count
//
// 图片总数查询
type TaobaoPicturePicturesCountAPIResponse struct {
	model.CommonResponse
	TaobaoPicturePicturesCountAPIResponseModel
}

// TaobaoPicturePicturesCountAPIResponseModel is 图片总数查询 成功返回结果
type TaobaoPicturePicturesCountAPIResponseModel struct {
	XMLName xml.Name `xml:"picture_pictures_count_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询的文件总数
	Totals int64 `json:"totals,omitempty" xml:"totals,omitempty"`
}
