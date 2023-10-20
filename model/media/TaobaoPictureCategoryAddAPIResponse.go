package media

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPictureCategoryAddAPIResponse 新增图片分类信息 API返回值
// taobao.picture.category.add
//
// 同一卖家最多添加500个图片分类，图片分类名称长度最大为20个字符
type TaobaoPictureCategoryAddAPIResponse struct {
	model.CommonResponse
	TaobaoPictureCategoryAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPictureCategoryAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPictureCategoryAddAPIResponseModel).Reset()
}

// TaobaoPictureCategoryAddAPIResponseModel is 新增图片分类信息 成功返回结果
type TaobaoPictureCategoryAddAPIResponseModel struct {
	XMLName xml.Name `xml:"picture_category_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 图片分类信息
	PictureCategory *PictureCategory `json:"picture_category,omitempty" xml:"picture_category,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPictureCategoryAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.PictureCategory = nil
}

var poolTaobaoPictureCategoryAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPictureCategoryAddAPIResponse)
	},
}

// GetTaobaoPictureCategoryAddAPIResponse 从 sync.Pool 获取 TaobaoPictureCategoryAddAPIResponse
func GetTaobaoPictureCategoryAddAPIResponse() *TaobaoPictureCategoryAddAPIResponse {
	return poolTaobaoPictureCategoryAddAPIResponse.Get().(*TaobaoPictureCategoryAddAPIResponse)
}

// ReleaseTaobaoPictureCategoryAddAPIResponse 将 TaobaoPictureCategoryAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPictureCategoryAddAPIResponse(v *TaobaoPictureCategoryAddAPIResponse) {
	v.Reset()
	poolTaobaoPictureCategoryAddAPIResponse.Put(v)
}
