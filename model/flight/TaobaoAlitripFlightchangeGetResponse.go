package flight

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取航变信息 APIResponse
taobao.alitrip.flightchange.get

查询航变是为了两个目的，阿里旅行抓取到未确认航变的航变信息源时可以由代理确认航变，同时对于确认航变的航变信息也共享给代理人做本体业务使用。
*/
type TaobaoAlitripFlightchangeGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alitrip_flightchange_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    ResultDO   *TaobaoAlitripFlightchangeGetResultDo `json:"result_d_o,omitempty" xml:"