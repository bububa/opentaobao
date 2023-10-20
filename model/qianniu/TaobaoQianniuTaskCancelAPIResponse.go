package qianniu

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqianniutaskcancelAPIResponse 取消轻任务 API返回值
// taobao.qianniu.task.cancel
//
// 由任务发起者调用
type TaobaoqianniutaskcancelAPIResponse struct {
	model.CommonResponse
	TaobaoqianniutaskcancelAPIResponseModel
}

// TaobaoqianniutaskcancelAPIResponseModel is 取消轻任务 成功返回结果
type TaobaoqianniutaskcancelAPIResponseModel struct {
	XMLName xml.Name `xml:"qianniu_task_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
