package media

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPictureCategoryGetAPIResponse 获取图片分类信息 API返回值
// taobao.picture.category.get
//
// 获取图片分类信息
type TaobaoPictureCategoryGetAPIResponse struct {
	model.CommonResponse
	TaobaoPictureCategoryGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPictureCategoryGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPictureCategoryGetAPIResponseModel).Reset()
}

// TaobaoPictureCategoryGetAPIResponseModel is 获取图片分类信息 成功返回结果
type TaobaoPictureCategoryGetAPIResponseModel struct {
	XMLName xml.Name `xml:"picture_category_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 图片分类
	PictureCategories []PictureCategory `json:"picture_categories,omitempty" xml:"picture_categories>picture_category,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPictureCategoryGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.PictureCategories = m.PictureCategories[:0]
}

var poolTaobaoPictureCategoryGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPictureCategoryGetAPIResponse)
	},
}

// GetTaobaoPictureCategoryGetAPIResponse 从 sync.Pool 获取 TaobaoPictureCategoryGetAPIResponse
func GetTaobaoPictureCategoryGetAPIResponse() *TaobaoPictureCategoryGetAPIResponse {
	return poolTaobaoPictureCategoryGetAPIResponse.Get().(*TaobaoPictureCategoryGetAPIResponse)
}

// ReleaseTaobaoPictureCategoryGetAPIResponse 将 TaobaoPictureCategoryGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPictureCategoryGetAPIResponse(v *TaobaoPictureCategoryGetAPIResponse) {
	v.Reset()
	poolTaobaoPictureCategoryGetAPIResponse.Put(v)
}
