package wangwang

import (
    "github.com/bububa/opentaobao/model"
)

/* 
客服评价详情接口 APIResponse
taobao.qianniu.kefueval.get

获取买家对客服的服务评价
*/
type TaobaoQianniuKefuevalGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQianniuKefuevalGetResponse `json:"qianniu_kefueval_get_response,omitempty"` 
    TaobaoQianniuKefuevalGetResponse
}

/* model for simplify = false
type TaobaoQianniuKefuevalGetResponse struct {

    // 评价结果数
    
    ResultCount   int64 `json:"result_count,omitempty"`
    

    // 评价明细
    
    StaffEvalDetails  struct {
        EvalDetail  []EvalDetail `json:"eval_detail,omitempty"`
    } `json:"staff_eval_details,omitempty"`
    

}
*/

type TaobaoQianniuKefuevalGetResponse struct {

    // 评价结果数
    ResultCount   int64 `json:"result_count,omitempty"`

    // 评价明细
    StaffEvalDetails   []EvalDetail `json:"staff_eval_details,omitempty"`

}
