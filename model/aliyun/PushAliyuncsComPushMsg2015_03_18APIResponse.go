package aliyun

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// PushAliyuncsComPushMsg2015_03_18APIResponse 消息推送 API返回值
// push.aliyuncs.com.pushMsg.2015-03-18
//
// 消息推送  ,支持指定用户/账号/广播等模式
type PushAliyuncsComPushMsg2015_03_18APIResponse struct {
	model.CommonResponse
	PushAliyuncsComPushMsg2015_03_18APIResponseModel
}

// PushAliyuncsComPushMsg2015_03_18APIResponseModel is 消息推送 成功返回结果
type PushAliyuncsComPushMsg2015_03_18APIResponseModel struct {
	XMLName xml.Name `xml:"push_aliyuncs_com_pushMsg_2015-03-18_response"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 该字段的值由服务端生成,返回给用户方便问题追查与定位。
	RequestId int64 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 消息ID
	ResponseParams int64 `json:"responseParams,omitempty" xml:"responseParams,omitempty"`
}
