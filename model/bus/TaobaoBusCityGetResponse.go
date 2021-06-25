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
    Response *TaobaoBusCityGetResponse `json:"taobao_bus_city_get_response,omitempty"`
}

type TaobaoBusCityGetResponse struct {

    // 城市返回结果
    Result   *CitySearchRp `json:"result,omitempty"`

}
