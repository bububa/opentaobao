package jstinteractive

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstinteractiveactivityqueryAPIResponse 互动任务活动查询接口 API返回值
// taobao.jst.interactive.activity.query
//
// 互动任务活动查询接口
type TaobaojstinteractiveactivityqueryAPIResponse struct {
	model.CommonResponse
	TaobaojstinteractiveactivityqueryAPIResponseModel
}

// TaobaojstinteractiveactivityqueryAPIResponseModel is 互动任务活动查询接口 成功返回结果
type TaobaojstinteractiveactivityqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_interactive_activity_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 活动列表，只返回有效的活动
	ActivityList []Activity `json:"activity_list,omitempty" xml:"activity_list>activity,omitempty"`
}
