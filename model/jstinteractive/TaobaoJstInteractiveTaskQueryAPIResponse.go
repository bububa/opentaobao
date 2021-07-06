package jstinteractive

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstInteractiveTaskQueryAPIResponse 互动任务列表查询接口 API返回值
// taobao.jst.interactive.task.query
//
// 查询用户的互动任务列表
type TaobaoJstInteractiveTaskQueryAPIResponse struct {
	model.CommonResponse
	TaobaoJstInteractiveTaskQueryAPIResponseModel
}

// TaobaoJstInteractiveTaskQueryAPIResponseModel is 互动任务列表查询接口 成功返回结果
type TaobaoJstInteractiveTaskQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_interactive_task_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果
	Data *InteractiveTaskQueryResponse `json:"data,omitempty" xml:"data,omitempty"`
	// 调用是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
