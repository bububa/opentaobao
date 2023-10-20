package jstinteractive

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstinteractiveactivityupdateAPIResponse 互动任务活动修改接口 API返回值
// taobao.jst.interactive.activity.update
//
// 互动任务活动修改接口
type TaobaojstinteractiveactivityupdateAPIResponse struct {
	model.CommonResponse
	TaobaojstinteractiveactivityupdateAPIResponseModel
}

// TaobaojstinteractiveactivityupdateAPIResponseModel is 互动任务活动修改接口 成功返回结果
type TaobaojstinteractiveactivityupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_interactive_activity_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 修改结果
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
