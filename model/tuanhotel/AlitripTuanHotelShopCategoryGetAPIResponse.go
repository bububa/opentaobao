package tuanhotel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTuanHotelShopCategoryGetAPIResponse 商家店铺类目查询 API返回值
// alitrip.tuan.hotel.shop.category.get
//
// 查询商家店铺类目信息
type AlitripTuanHotelShopCategoryGetAPIResponse struct {
	model.CommonResponse
	AlitripTuanHotelShopCategoryGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripTuanHotelShopCategoryGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripTuanHotelShopCategoryGetAPIResponseModel).Reset()
}

// AlitripTuanHotelShopCategoryGetAPIResponseModel is 商家店铺类目查询 成功返回结果
type AlitripTuanHotelShopCategoryGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_tuan_hotel_shop_category_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 一级类目列表
	TopRootShopCategoryList []TopRootShopCategoryVoList `json:"top_root_shop_category_list,omitempty" xml:"top_root_shop_category_list>top_root_shop_category_vo_list,omitempty"`
	// code
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Status bool `json:"status,omitempty" xml:"status,omitempty"`
}

// Reset 清空结构体
func (m *AlitripTuanHotelShopCategoryGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TopRootShopCategoryList = m.TopRootShopCategoryList[:0]
	m.MsgCode = ""
	m.Message = ""
	m.Status = false
}

var poolAlitripTuanHotelShopCategoryGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripTuanHotelShopCategoryGetAPIResponse)
	},
}

// GetAlitripTuanHotelShopCategoryGetAPIResponse 从 sync.Pool 获取 AlitripTuanHotelShopCategoryGetAPIResponse
func GetAlitripTuanHotelShopCategoryGetAPIResponse() *AlitripTuanHotelShopCategoryGetAPIResponse {
	return poolAlitripTuanHotelShopCategoryGetAPIResponse.Get().(*AlitripTuanHotelShopCategoryGetAPIResponse)
}

// ReleaseAlitripTuanHotelShopCategoryGetAPIResponse 将 AlitripTuanHotelShopCategoryGetAPIResponse 保存到 sync.Pool
func ReleaseAlitripTuanHotelShopCategoryGetAPIResponse(v *AlitripTuanHotelShopCategoryGetAPIResponse) {
	v.Reset()
	poolAlitripTuanHotelShopCategoryGetAPIResponse.Put(v)
}
