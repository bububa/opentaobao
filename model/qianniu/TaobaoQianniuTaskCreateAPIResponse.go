package qianniu

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqianniutaskcreateAPIResponse 创建轻任务 API返回值
// taobao.qianniu.task.create
//
// 发起一个轻任务，分配给多个执行者，并发送消息提醒，由任务发起者调用
type TaobaoqianniutaskcreateAPIResponse struct {
	model.CommonResponse
	TaobaoqianniutaskcreateAPIResponseModel
}

// TaobaoqianniutaskcreateAPIResponseModel is 创建轻任务 成功返回结果
type TaobaoqianniutaskcreateAPIResponseModel struct {
	XMLName xml.Name `xml:"qianniu_task_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创建的任务元数据
	Result *QtaskMetadata `json:"result,omitempty" xml:"result,omitempty"`
}
