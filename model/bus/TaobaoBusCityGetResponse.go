package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
城市接口 APIResponse
taobao.bus.city.get

汽车票出发城市获取接口，获取所有出发城市
*/
type TaobaoBusCityGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"bus_city_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 城市返回结果
    
    Result   *CitySearchRp `json:"result,omitempty" xml:"