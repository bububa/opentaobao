package itpolicy

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripItFareQuerytaskAPIResponse 【国际机票自有政策】批量操作结果查询 API返回值
// taobao.alitrip.it.fare.querytask
//
// 批量操作同步返回任务id，后台生成异步任务，通过此接口查询批量操作的执行结果
type TaobaoAlitripItFareQuerytaskAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripItFareQuerytaskAPIResponseModel
}

// TaobaoAlitripItFareQuerytaskAPIResponseModel is 【国际机票自有政策】批量操作结果查询 成功返回结果
type TaobaoAlitripItFareQuerytaskAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_it_fare_querytask_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 任务执行失败，会返回一条失败信息。如果是导入任务，会返回每条导入失败的政策说明，最多只返回200条失败信息。
	Errors []ErrorFareRow `json:"errors,omitempty" xml:"errors>error_fare_row,omitempty"`
	// json格式的字符串，扩展属性，预留
	ExtendAttributes string `json:"extend_attributes,omitempty" xml:"extend_attributes,omitempty"`
	// 完成时间
	GmtFinished string `json:"gmt_finished,omitempty" xml:"gmt_finished,omitempty"`
	// 成功处理条数
	ProcessAmount int64 `json:"process_amount,omitempty" xml:"process_amount,omitempty"`
	// 任务状态，1 处理中，2 处理失败，3 处理完毕
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
