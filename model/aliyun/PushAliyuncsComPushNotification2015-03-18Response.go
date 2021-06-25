package aliyun

import (
    "github.com/bububa/opentaobao/model"
)

/* 
推送通知 APIResponse
push.aliyuncs.com.pushNotification.2015-03-18

pushNotification
*/
type PushAliyuncsComPushNotification2015-03-18APIResponse struct {
    model.CommonResponse
    Response *PushAliyuncsComPushNotification2015-03-18Response `json:"push_aliyuncs_com_pushNotification_2015-03-18_response,omitempty"`
}

type PushAliyuncsComPushNotification2015-03-18Response struct {

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 该字段的值由服务端生成,返回给用户方便问题追查与定位。
    RequestId   int64 `json:"requestId,omitempty"`

    // 消息ID,用于查询
    ResponseParams   int64 `json:"responseParams,omitempty"`

}
