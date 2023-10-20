package media

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPictureCategoryUpdateAPIResponse 更新图片分类 API返回值
// taobao.picture.category.update
//
// 更新图片分类的名字，或者更新图片分类的父分类（即分类移动）。只能移动2级分类到非2级分类，默认分类和1级分类不可移动。
type TaobaoPictureCategoryUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoPictureCategoryUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPictureCategoryUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPictureCategoryUpdateAPIResponseModel).Reset()
}

// TaobaoPictureCategoryUpdateAPIResponseModel is 更新图片分类 成功返回结果
type TaobaoPictureCategoryUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"picture_category_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 更新图片分类是否成功
	Done bool `json:"done,omitempty" xml:"done,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPictureCategoryUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Done = false
}

var poolTaobaoPictureCategoryUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPictureCategoryUpdateAPIResponse)
	},
}

// GetTaobaoPictureCategoryUpdateAPIResponse 从 sync.Pool 获取 TaobaoPictureCategoryUpdateAPIResponse
func GetTaobaoPictureCategoryUpdateAPIResponse() *TaobaoPictureCategoryUpdateAPIResponse {
	return poolTaobaoPictureCategoryUpdateAPIResponse.Get().(*TaobaoPictureCategoryUpdateAPIResponse)
}

// ReleaseTaobaoPictureCategoryUpdateAPIResponse 将 TaobaoPictureCategoryUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPictureCategoryUpdateAPIResponse(v *TaobaoPictureCategoryUpdateAPIResponse) {
	v.Reset()
	poolTaobaoPictureCategoryUpdateAPIResponse.Put(v)
}
