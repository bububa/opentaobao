package taotv

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotaotvcarouselcategorylistAPIResponse 获取轮播分类列表 API返回值
// taobao.taotv.carousel.category.list
//
// 获取轮播分类列表
type TaobaotaotvcarouselcategorylistAPIResponse struct {
	model.CommonResponse
	TaobaotaotvcarouselcategorylistAPIResponseModel
}

// TaobaotaotvcarouselcategorylistAPIResponseModel is 获取轮播分类列表 成功返回结果
type TaobaotaotvcarouselcategorylistAPIResponseModel struct {
	XMLName xml.Name `xml:"taotv_carousel_category_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaotaotvcarouselcategorylistResult `json:"result,omitempty" xml:"result,omitempty"`
}
