package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
线下自助机通知出票接口 API返回值 
taobao.bus.tvmbookorder.set

出票，当成功的时候告知出票；当失败的时候告知出票失败，飞猪退款给用户。
*/
type TaobaoBusTvmbookorderSetAPIResponse struct {
    model.CommonResponse
    TaobaoBusTvmbookorderSetResponse
}

// 线下自助机通知出票接口 成功返回结果
type TaobaoBusTvmbookorderSetResponse struct {
    XMLName xml.Name `xml:"bus_tvmbookorder_set_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // errorCode
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // errorMsg
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
    // success
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
