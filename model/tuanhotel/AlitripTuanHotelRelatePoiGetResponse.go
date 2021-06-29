package tuanhotel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店团购关联POI APIResponse
alitrip.tuan.hotel.relate.poi.get

查询酒店团购关联POI
*/
type AlitripTuanHotelRelatePoiGetAPIResponse struct {
    model.CommonResponse
    AlitripTuanHotelRelatePoiGetResponse
}

type AlitripTuanHotelRelatePoiGetResponse struct {
    XMLName xml.Name `xml:"alitrip_tuan_hotel_relate_poi_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 关联POI列表
    
    RelatedPoiDetailList   []RelatedPoiDetailVoList `json:"related_poi_detail_list,omitempty" xml:"related_poi_detail_list>related_poi_detail_vo_list,omitempty"`
    
    
    // 错误码
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 操作状态
    
    Status   bool `json:"status,omitempty" xml:"status,omitempty"`

    
    // 错误信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
}
