package qianniu

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQianniuTasksGetAPIResponse 获取指定的任务 API返回值
// taobao.qianniu.tasks.get
//
// 获取指定的任务，可用的参数组合：&lt;br/&gt;task_ids + need_meta + fields：精确查找&lt;br/&gt;biz_type + sub_biz_type + biz_ids + need_meta + fields：按照业务ID查找&lt;br/&gt;biz_type + sub_biz_type + sender_uid + need_meta + fields：按照发起者查找&lt;br/&gt;biz_type + sub_biz_type + receiver_uid + need_meta + fields：按照执行者查找&lt;br/&gt;biz_type+modify_start_time+modify_end_time+fields:能支持指定修改时间的查询，用于增量查询等
type TaobaoQianniuTasksGetAPIResponse struct {
	model.CommonResponse
	TaobaoQianniuTasksGetAPIResponseModel
}

// TaobaoQianniuTasksGetAPIResponseModel is 获取指定的任务 成功返回结果
type TaobaoQianniuTasksGetAPIResponseModel struct {
	XMLName xml.Name `xml:"qianniu_tasks_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的任务列表
	Tasks []Qtask `json:"tasks,omitempty" xml:"tasks>qtask,omitempty"`
}
