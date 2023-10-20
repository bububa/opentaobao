package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtriphotelordersearchAPIResponse 搜索酒店订单列表 API返回值
// alitrip.btrip.hotel.order.search
//
// 企业获取商旅酒店订单数据
type AlitripbtriphotelordersearchAPIResponse struct {
	model.CommonResponse
	AlitripbtriphotelordersearchAPIResponseModel
}

// AlitripbtriphotelordersearchAPIResponseModel is 搜索酒店订单列表 成功返回结果
type AlitripbtriphotelordersearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_hotel_order_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *BtriphomeResultSupport `json:"result,omitempty" xml:"result,omitempty"`
}
