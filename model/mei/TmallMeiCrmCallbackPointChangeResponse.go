package mei

import (
    "github.com/bububa/opentaobao/model"
)

/* 
品牌积分变更回调API APIResponse
tmall.mei.crm.callback.point.change

线下品牌积分变更消息回调API，告诉积分扣减或者累加是否成功。
*/
type TmallMeiCrmCallbackPointChangeAPIResponse struct {
    model.CommonResponse
    // Response *TmallMeiCrmCallbackPointChangeResponse `json:"tmall_mei_crm_callback_point_change_response,omitempty"` 
    TmallMeiCrmCallbackPointChangeResponse
}

/* model for simplify = false
type TmallMeiCrmCallbackPointChangeResponse struct {

    // 是否成功
    
    ResultSuccess   bool `json:"result_success,omitempty"`
    

}
*/

type TmallMeiCrmCallbackPointChangeResponse struct {

    // 是否成功
    ResultSuccess   bool `json:"result_success,omitempty"`

}
