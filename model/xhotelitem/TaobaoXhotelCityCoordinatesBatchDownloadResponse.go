package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
下载飞猪国际城市结果 APIResponse
taobao.xhotel.city.coordinates.batch.download

给国际酒店供应商提供计算对应飞猪城市的服务，免去城市名称匹配流程，加快对接流程
*/
type TaobaoXhotelCityCoordinatesBatchDownloadAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelCityCoordinatesBatchDownloadResponse
}

type TaobaoXhotelCityCoordinatesBatchDownloadResponse struct {
    XMLName xml.Name `xml:"xhotel_city_coordinates_batch_download_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 经纬度计算结果列表
    
    CoordinateList   []Coordinate `json:"coordinate_list,omitempty" xml:"coordinate_list>coordinate,omitempty"`
    
    
}
