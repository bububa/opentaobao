package tmallsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建主动预约开通条件 API返回值 
tmall.servicecenter.reservecond.create

1、设置主动预约开通条件
*/
type TmallServicecenterReservecondCreateAPIResponse struct {
    model.CommonResponse
    TmallServicecenterReservecondCreateAPIResponseModel
}

// 创建主动预约开通条件 成功返回结果
type TmallServicecenterReservecondCreateAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_servicecenter_reservecond_create_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 是否成功
    MsgSuccess   bool `json:"msg_success,omitempty" xml:"msg_success,omitempty"`
    // 返回信息
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // 返回code
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
}
