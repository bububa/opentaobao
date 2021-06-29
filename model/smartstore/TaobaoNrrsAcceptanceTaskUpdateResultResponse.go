package smartstore

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新验收任务支付宝结果 API返回值 
taobao.nrrs.acceptance.task.updateResult

智慧门店商家验收任务检查相关接口-更新支付宝的验收结果。
*/
type TaobaoNrrsAcceptanceTaskUpdateResultAPIResponse struct {
    model.CommonResponse
    TaobaoNrrsAcceptanceTaskUpdateResultResponse
}

// 更新验收任务支付宝结果 成功返回结果
type TaobaoNrrsAcceptanceTaskUpdateResultResponse struct {
    XMLName xml.Name `xml:"nrrs_acceptance_task_updateResult_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
