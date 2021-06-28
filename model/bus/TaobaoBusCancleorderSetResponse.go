package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取消订单 APIResponse
taobao.bus.cancleorder.set

取消订单
*/
type TaobaoBusCancleorderSetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"bus_cancleorder_set_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误代码
    
    ErrorCode1   string `json:"error_code1,omitempty" xml:"