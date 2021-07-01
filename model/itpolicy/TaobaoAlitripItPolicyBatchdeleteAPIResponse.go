package itpolicy

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票销售规则】批量删除 API返回值 
taobao.alitrip.it.policy.batchdelete

批量删除销售规则，单次删除最大5万，大于5万时候提示失败，需要缩小删除条件。此接口同步返回任务id，异步执行操作。每个接入方最多同时只能有10个处理中的任务，超过后直接返回失败。
*/
type TaobaoAlitripItPolicyBatchdeleteAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripItPolicyBatchdeleteAPIResponseModel
}

// 【国际机票销售规则】批量删除 成功返回结果
type TaobaoAlitripItPolicyBatchdeleteAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_it_policy_batchdelete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 任务id
    TaskId   int64 `json:"task_id,omitempty" xml:"task_id,omitempty"`
    // 扩展字段
    ExtendAttributes   string `json:"extend_attributes,omitempty" xml:"extend_attributes,omitempty"`
}
