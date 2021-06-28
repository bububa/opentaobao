package aliyun

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
推送通知 APIResponse
push.aliyuncs.com.pushNotification.2015-03-18

pushNotification
*/
type PushAliyuncsComPushNotification2015-03-18APIResponse struct {
    model.CommonResponse
    PushAliyuncsComPushNotification2015-03-18Response
}

type PushAliyuncsComPushNotification2015-03-18Response struct {
    XMLName xml.Name `xml:"push_aliyuncs_com_pushNotification_2015-03-18_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`

    
    // 该字段的值由服务端生成,返回给用户方便问题追查与定位。
    
    RequestId   int64 `json:"requestId,omitempty" xml:"requestId,omitempty"`

    
    // 消息ID,用于查询
    
    ResponseParams   int64 `json:"responseParams,omitempty" xml:"responseParams,omitempty"`

    
}
