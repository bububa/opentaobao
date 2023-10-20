package hotelhstdf

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitriphotelhstdfhotelroomstaticgetAPIResponse 根据类型查询静态字段 API返回值
// alitrip.hotel.hstdf.hotelroomstatic.get
//
// 根据类型查询分页静态字段
type AlitriphotelhstdfhotelroomstaticgetAPIResponse struct {
	model.CommonResponse
	AlitriphotelhstdfhotelroomstaticgetAPIResponseModel
}

// AlitriphotelhstdfhotelroomstaticgetAPIResponseModel is 根据类型查询静态字段 成功返回结果
type AlitriphotelhstdfhotelroomstaticgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_hotel_hstdf_hotelroomstatic_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// top返回结果
	Result *TopResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
