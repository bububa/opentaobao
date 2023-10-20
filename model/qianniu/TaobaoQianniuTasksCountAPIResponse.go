package qianniu

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqianniutaskscountAPIResponse 任务查询条数接口 API返回值
// taobao.qianniu.tasks.count
//
// 任务查询条数接口
type TaobaoqianniutaskscountAPIResponse struct {
	model.CommonResponse
	TaobaoqianniutaskscountAPIResponseModel
}

// TaobaoqianniutaskscountAPIResponseModel is 任务查询条数接口 成功返回结果
type TaobaoqianniutaskscountAPIResponseModel struct {
	XMLName xml.Name `xml:"qianniu_tasks_count_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 符合查询条件的总条数
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}
