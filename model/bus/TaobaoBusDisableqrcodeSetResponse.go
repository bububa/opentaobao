package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
自助机失效二维码 APIResponse
taobao.bus.disableqrcode.set

使创建的二维码失效
*/
type TaobaoBusDisableqrcodeSetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"bus_disableqrcode_set_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // errorCode
    
    ResultCode   string `json:"result_code,omitempty" xml:"