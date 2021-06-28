package qianniu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
任务查询条数接口 APIResponse
taobao.qianniu.tasks.count

任务查询条数接口
*/
type TaobaoQianniuTasksCountAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qianniu_tasks_count_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 符合查询条件的总条数
    
    Result   int64 `json:"result,omitempty" xml:"