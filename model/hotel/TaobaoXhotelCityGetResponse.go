package hotel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店城市数据获取接口 APIResponse
taobao.xhotel.city.get

引流API，对外提供酒店城市数据
*/
type TaobaoXhotelCityGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"xhotel_city_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 总数
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"