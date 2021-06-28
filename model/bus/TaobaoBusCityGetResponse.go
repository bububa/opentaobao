package bus

import (
    "github.com/bububa/opentaobao/model"
)

/* 
城市接口 APIResponse
taobao.bus.city.get

汽车票出发城市获取接口，获取所有出发城市
*/
type TaobaoBusCityGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBusCityGetResponse `json:"bus_city_get_response,omitempty"` 
    TaobaoBusCityGetResponse
}

/* model for simplify = false
type TaobaoBusCityGetResponse struct {

    // 城市返回结果
    
    Result  *struct {
        CitySearchRp  *CitySearchRp `json:"city_search_rp,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoBusCityGetResponse struct {

    // 城市返回结果
    Result   *CitySearchRp `json:"result,omitempty"`

}
