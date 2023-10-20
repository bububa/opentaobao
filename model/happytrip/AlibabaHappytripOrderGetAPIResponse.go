package happytrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabahappytripordergetAPIResponse 获取欢行统一订单模型 API返回值
// alibaba.happytrip.order.get
//
// 通过订单id获取欢行统一订单模型数据
type AlibabahappytripordergetAPIResponse struct {
	model.CommonResponse
	AlibabahappytripordergetAPIResponseModel
}

// AlibabahappytripordergetAPIResponseModel is 获取欢行统一订单模型 成功返回结果
type AlibabahappytripordergetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_happytrip_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单联系人列表
	Contacts []ContactDto `json:"contacts,omitempty" xml:"contacts>contact_dto,omitempty"`
	// 机票预订信息列表
	HotelBooks []HotelBookDto `json:"hotel_books,omitempty" xml:"hotel_books>hotel_book_dto,omitempty"`
	// 酒店资源详情
	HotelResources []ResourceHotelDto `json:"hotel_resources,omitempty" xml:"hotel_resources>resource_hotel_dto,omitempty"`
	// 订单资源列表
	Resources []ResourceMainDto `json:"resources,omitempty" xml:"resources>resource_main_dto,omitempty"`
	// 订单出行人列表
	Tourists []TouristDto `json:"tourists,omitempty" xml:"tourists>tourist_dto,omitempty"`
	// 订单扩展信息
	OrderExtendsInfo *OrderExtendsDto `json:"order_extends_info,omitempty" xml:"order_extends_info,omitempty"`
	// 订单基本信息
	OrderInfo *OrderDto `json:"order_info,omitempty" xml:"order_info,omitempty"`
}
