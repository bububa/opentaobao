package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵内容库视频分集数据推送接口 API返回值 
alibaba.ailabs.aligenie.openvideo.push

天猫精灵内容库视频分集数据推送接口
*/
type AlibabaAilabsAligenieOpenvideoPushAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsAligenieOpenvideoPushResponse
}

// 天猫精灵内容库视频分集数据推送接口 成功返回结果
type AlibabaAilabsAligenieOpenvideoPushResponse struct {
    XMLName xml.Name `xml:"alibaba_ailabs_aligenie_openvideo_push_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 状态码
    RetCode   int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
    // 描述
    RetMsg   string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
}
