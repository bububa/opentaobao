package smartstore

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新验收任务支付宝结果 APIResponse
taobao.nrrs.acceptance.task.updateResult

智慧门店商家验收任务检查相关接口-更新支付宝的验收结果。
*/
type TaobaoNrrsAcceptanceTaskUpdateResultAPIResponse struct {
    model.CommonResponse
    TaobaoNrrsAcceptanceTaskUpdateResultResponse
}

type TaobaoNrrsAcceptanceTaskUpdateResultResponse struct {
    XMLName xml.Name `xml:"nrrs_acceptance_task_updateResult_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
