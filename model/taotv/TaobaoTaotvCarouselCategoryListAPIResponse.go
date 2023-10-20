package taotv

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTaotvCarouselCategoryListAPIResponse 获取轮播分类列表 API返回值
// taobao.taotv.carousel.category.list
//
// 获取轮播分类列表
type TaobaoTaotvCarouselCategoryListAPIResponse struct {
	model.CommonResponse
	TaobaoTaotvCarouselCategoryListAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTaotvCarouselCategoryListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTaotvCarouselCategoryListAPIResponseModel).Reset()
}

// TaobaoTaotvCarouselCategoryListAPIResponseModel is 获取轮播分类列表 成功返回结果
type TaobaoTaotvCarouselCategoryListAPIResponseModel struct {
	XMLName xml.Name `xml:"taotv_carousel_category_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoTaotvCarouselCategoryListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTaotvCarouselCategoryListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTaotvCarouselCategoryListAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTaotvCarouselCategoryListAPIResponse)
	},
}

// GetTaobaoTaotvCarouselCategoryListAPIResponse 从 sync.Pool 获取 TaobaoTaotvCarouselCategoryListAPIResponse
func GetTaobaoTaotvCarouselCategoryListAPIResponse() *TaobaoTaotvCarouselCategoryListAPIResponse {
	return poolTaobaoTaotvCarouselCategoryListAPIResponse.Get().(*TaobaoTaotvCarouselCategoryListAPIResponse)
}

// ReleaseTaobaoTaotvCarouselCategoryListAPIResponse 将 TaobaoTaotvCarouselCategoryListAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTaotvCarouselCategoryListAPIResponse(v *TaobaoTaotvCarouselCategoryListAPIResponse) {
	v.Reset()
	poolTaobaoTaotvCarouselCategoryListAPIResponse.Put(v)
}
