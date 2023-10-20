package aliyun

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// PushaliyuncscompushMsg20150318APIResponse 消息推送 API返回值
// push.aliyuncs.com.pushMsg.2015-03-18
//
// 消息推送  ,支持指定用户/账号/广播等模式
type PushaliyuncscompushMsg20150318APIResponse struct {
	model.CommonResponse
	PushaliyuncscompushMsg20150318APIResponseModel
}

// PushaliyuncscompushMsg20150318APIResponseModel is 消息推送 成功返回结果
type PushaliyuncscompushMsg20150318APIResponseModel struct {
	XMLName xml.Name `xml:"push_aliyuncs_com_pushMsg_2015-03-18_response"`
	// 该字段的值由服务端生成,返回给用户方便问题追查与定位。
	RequestId int64 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 消息ID
	ResponseParams int64 `json:"responseParams,omitempty" xml:"responseParams,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
