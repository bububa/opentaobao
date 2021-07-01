package alisports

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
乐动力线-微信设备绑定通知 API返回值 
alibaba.alisports.ldlserver.wxevent.post

转发弹外微信服务到弹内
*/
type AlibabaAlisportsLdlserverWxeventPostAPIResponse struct {
    model.CommonResponse
    AlibabaAlisportsLdlserverWxeventPostAPIResponseModel
}

// 乐动力线-微信设备绑定通知 成功返回结果
type AlibabaAlisportsLdlserverWxeventPostAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alisports_ldlserver_wxevent_post_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 业务成功
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // 错误码
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
}
