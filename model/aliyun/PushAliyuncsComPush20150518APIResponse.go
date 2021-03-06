package aliyun

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// PushAliyuncsComPush20150518APIResponse 云推送指令API API返回值
// push.aliyuncs.com.push.20150518
//
// 阿里云推送新增API，允许一条推送指令同时发布到多个终端上。
type PushAliyuncsComPush20150518APIResponse struct {
	model.CommonResponse
	PushAliyuncsComPush20150518APIResponseModel
}

// PushAliyuncsComPush20150518APIResponseModel is 云推送指令API 成功返回结果
type PushAliyuncsComPush20150518APIResponseModel struct {
	XMLName xml.Name `xml:"push_aliyuncs_com_push_20150518_response"`
	// 消息ID,用于查询
	ResponseParams string `json:"responseParams,omitempty" xml:"responseParams,omitempty"`
	// 该字段的值由服务端生成,返回给用户方便问题追查与定位。
	RequestId int64 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
