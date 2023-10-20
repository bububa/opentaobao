package axindata

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelaxinhotelpricebatchgetAPIResponse 阿信酒店批量报价查询接口 API返回值
// taobao.alitrip.travel.axin.hotel.price.batch.get
//
// 阿信酒店批量报价查询接口
type TaobaoalitriptravelaxinhotelpricebatchgetAPIResponse struct {
	model.CommonResponse
	TaobaoalitriptravelaxinhotelpricebatchgetAPIResponseModel
}

// TaobaoalitriptravelaxinhotelpricebatchgetAPIResponseModel is 阿信酒店批量报价查询接口 成功返回结果
type TaobaoalitriptravelaxinhotelpricebatchgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotel_price_batch_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *BaseResultApiDto `json:"result,omitempty" xml:"result,omitempty"`
}
