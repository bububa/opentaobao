package tmallsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
买家拒收 API返回值 
tmall.servicecenter.workcard.refuse

买家拒收通知接口
*/
type TmallServicecenterWorkcardRefuseAPIResponse struct {
    model.CommonResponse
    TmallServicecenterWorkcardRefuseResponse
}

// 买家拒收 成功返回结果
type TmallServicecenterWorkcardRefuseResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_workcard_refuse_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回信息
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // 返回code
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    // 是否成功
    MsgSuccess   bool `json:"msg_success,omitempty" xml:"msg_success,omitempty"`
}
