package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
线下自助机通知出票接口 APIResponse
taobao.bus.tvmbookorder.set

出票，当成功的时候告知出票；当失败的时候告知出票失败，飞猪退款给用户。
*/
type TaobaoBusTvmbookorderSetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"bus_tvmbookorder_set_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // errorCode
    
    ResultCode   string `json:"result_code,omitempty" xml:"