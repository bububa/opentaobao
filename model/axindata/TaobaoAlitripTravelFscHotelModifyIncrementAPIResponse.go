package axindata

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscHotelModifyIncrementAPIResponse 酒店价库变更列表查询-供销平台 API返回值
// taobao.alitrip.travel.fsc.hotel.modify.increment
//
// 按照时间纬度查询酒店变更列表
type TaobaoAlitripTravelFscHotelModifyIncrementAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelFscHotelModifyIncrementAPIResponseModel
}

// TaobaoAlitripTravelFscHotelModifyIncrementAPIResponseModel is 酒店价库变更列表查询-供销平台 成功返回结果
type TaobaoAlitripTravelFscHotelModifyIncrementAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_hotel_modify_increment_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口应答对象
	Result *BaseResultApiDto `json:"result,omitempty" xml:"result,omitempty"`
}
