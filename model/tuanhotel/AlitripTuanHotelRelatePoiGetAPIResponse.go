package tuanhotel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTuanHotelRelatePoiGetAPIResponse
酒店团购关联POI API返回值
alitrip.tuan.hotel.relate.poi.get

查询酒店团购关联POI */
type AlitripTuanHotelRelatePoiGetAPIResponse struct {
	model.CommonResponse
	AlitripTuanHotelRelatePoiGetAPIResponseModel
}

// AlitripTuanHotelRelatePoiGetAPIResponseModel is 酒店团购关联POI 成功返回结果
type AlitripTuanHotelRelatePoiGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_tuan_hotel_relate_poi_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 关联POI列表
	RelatedPoiDetailList []RelatedPoiDetailVoList `json:"related_poi_detail_list,omitempty" xml:"related_poi_detail_list>related_poi_detail_vo_list,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 操作状态
	Status bool `json:"status,omitempty" xml:"status,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
