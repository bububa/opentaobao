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
    TaobaoBusCancleorderSetResponse
}

type TaobaoBusCancleorderSetResponse struct {
    XMLName xml.Name `xml:"bus_cancleorder_set_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误代码
    
    ErrorCode1   string `json:"error_code1,omitempty" xml:"error_code1,omitempty"`

    
    // 错误描述
    
    ErrorMsg1   string `json:"error_msg1,omitempty" xml:"error_msg1,omitempty"`

    
    // success
    
    Success1   bool `json:"success1,omitempty" xml:"success1,omitempty"`

    
}
