package itpolicy

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripitfarebatchdeleteAPIResponse 【国际机票自有政策】批量删除 API返回值
// taobao.alitrip.it.fare.batchdelete
//
// 批量删除自有政策，单次删除最大5万，大于5万时候提示失败，需要缩小删除条件。此接口同步返回任务id，异步执行操作。每个接入方最多同时只能有10个处理中的任务，超过后直接返回失败。
type TaobaoalitripitfarebatchdeleteAPIResponse struct {
	model.CommonResponse
	TaobaoalitripitfarebatchdeleteAPIResponseModel
}

// TaobaoalitripitfarebatchdeleteAPIResponseModel is 【国际机票自有政策】批量删除 成功返回结果
type TaobaoalitripitfarebatchdeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_it_fare_batchdelete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// json格式的字符串，扩展属性，预留
	ExtendAttributes string `json:"extend_attributes,omitempty" xml:"extend_attributes,omitempty"`
	// 任务id，可以根据任务id调用querytask查询执行结果
	TaskId int64 `json:"task_id,omitempty" xml:"task_id,omitempty"`
}
