package smartstore

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新验收任务支付宝结果 APIRequest
taobao.nrrs.acceptance.task.updateResult

智慧门店商家验收任务检查相关接口-更新支付宝的验收结果。
*/
type TaobaoNrrsAcceptanceTaskUpdateResultRequest struct {
    model.Params

    // 任务ID
    taskId   string 

    // 系统自动生成
    alipayResultList   []AlipayCheckResult 

}

func NewTaobaoNrrsAcceptanceTaskUpdateResultRequest() *TaobaoNrrsAcceptanceTaskUpdateResultRequest{
    return &TaobaoNrrsAcceptanceTaskUpdateResultRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoNrrsAcceptanceTaskUpdateResultRequest) GetApiMethodName() string {
    return "taobao.nrrs.acceptance.task.updateResult"
}

func (r TaobaoNrrsAcceptanceTaskUpdateResultRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoNrrsAcceptanceTaskUpdateResultRequest) SetTaskId(taskId string) error {
    r.taskId = taskId
    r.Set("task_id", taskId)
    return nil
}

func (r TaobaoNrrsAcceptanceTaskUpdateResultRequest) GetTaskId() string {
    return r.taskId
}

func (r *TaobaoNrrsAcceptanceTaskUpdateResultRequest) SetAlipayResultList(alipayResultList []AlipayCheckResult) error {
    r.alipayResultList = alipayResultList
    r.Set("alipay_result_list", alipayResultList)
    return nil
}

func (r TaobaoNrrsAcceptanceTaskUpdateResultRequest) GetAlipayResultList() []AlipayCheckResult {
    return r.alipayResultList
}

