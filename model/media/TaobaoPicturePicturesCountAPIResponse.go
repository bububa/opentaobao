package media

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPicturePicturesCountAPIResponse 图片总数查询 API返回值
// taobao.picture.pictures.count
//
// 图片总数查询，目前出于对数据库的保护暂不支持此功能
type TaobaoPicturePicturesCountAPIResponse struct {
	model.CommonResponse
	TaobaoPicturePicturesCountAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPicturePicturesCountAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPicturePicturesCountAPIResponseModel).Reset()
}

// TaobaoPicturePicturesCountAPIResponseModel is 图片总数查询 成功返回结果
type TaobaoPicturePicturesCountAPIResponseModel struct {
	XMLName xml.Name `xml:"picture_pictures_count_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询的文件总数
	Totals int64 `json:"totals,omitempty" xml:"totals,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPicturePicturesCountAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Totals = 0
}

var poolTaobaoPicturePicturesCountAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPicturePicturesCountAPIResponse)
	},
}

// GetTaobaoPicturePicturesCountAPIResponse 从 sync.Pool 获取 TaobaoPicturePicturesCountAPIResponse
func GetTaobaoPicturePicturesCountAPIResponse() *TaobaoPicturePicturesCountAPIResponse {
	return poolTaobaoPicturePicturesCountAPIResponse.Get().(*TaobaoPicturePicturesCountAPIResponse)
}

// ReleaseTaobaoPicturePicturesCountAPIResponse 将 TaobaoPicturePicturesCountAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPicturePicturesCountAPIResponse(v *TaobaoPicturePicturesCountAPIResponse) {
	v.Reset()
	poolTaobaoPicturePicturesCountAPIResponse.Put(v)
}
