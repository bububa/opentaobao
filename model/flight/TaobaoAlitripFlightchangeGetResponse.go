package flight

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取航变信息 APIResponse
taobao.alitrip.flightchange.get

查询航变是为了两个目的，阿里旅行抓取到未确认航变的航变信息源时可以由代理确认航变，同时对于确认航变的航变信息也共享给代理人做本体业务使用。
*/
type TaobaoAlitripFlightchangeGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAlitripFlightchangeGetResponse `json:"alitrip_flightchange_get_response,omitempty"` 
    TaobaoAlitripFlightchangeGetResponse
}

/* model for simplify = false
type TaobaoAlitripFlightchangeGetResponse struct {

    // result
    
    ResultDO  *struct {
        TaobaoAlitripFlightchangeGetResultDo  *TaobaoAlitripFlightchangeGetResultDo `json:"taobao_alitrip_flightchange_get_result_do,omitempty"`
    } `json:"result_d_o,omitempty"`
    

}
*/

type TaobaoAlitripFlightchangeGetResponse struct {

    // result
    ResultDO   *TaobaoAlitripFlightchangeGetResultDo `json:"result_d_o,omitempty"`

}
