package qianniu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
完成轻任务 API返回值 
taobao.qianniu.task.finish

由任务执行者调用
*/
type TaobaoQianniuTaskFinishAPIResponse struct {
    model.CommonResponse
    TaobaoQianniuTaskFinishResponse
}

// 完成轻任务 成功返回结果
type TaobaoQianniuTaskFinishResponse struct {
    XMLName xml.Name `xml:"qianniu_task_finish_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 是否成功
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`
}
