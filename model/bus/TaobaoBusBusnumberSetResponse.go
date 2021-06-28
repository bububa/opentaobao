package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家汽车票车次更新通知接口 APIResponse
taobao.bus.busnumber.set

商家汽车票车次更新后，调用该接口通知平台。
*/
type TaobaoBusBusnumberSetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"bus_busnumber_set_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误描述
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"