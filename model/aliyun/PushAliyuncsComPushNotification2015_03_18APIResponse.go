package aliyun

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// PushAliyuncsComPushNotification2015_03_18APIResponse 推送通知 API返回值
// push.aliyuncs.com.pushNotification.2015-03-18
//
// pushNotification
type PushAliyuncsComPushNotification2015_03_18APIResponse struct {
	model.CommonResponse
	PushAliyuncsComPushNotification2015_03_18APIResponseModel
}

// PushAliyuncsComPushNotification2015_03_18APIResponseModel is 推送通知 成功返回结果
type PushAliyuncsComPushNotification2015_03_18APIResponseModel struct {
	XMLName xml.Name `xml:"push_aliyuncs_com_pushNotification_2015-03-18_response"`
	// 该字段的值由服务端生成,返回给用户方便问题追查与定位。
	RequestId int64 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 消息ID,用于查询
	ResponseParams int64 `json:"responseParams,omitempty" xml:"responseParams,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
