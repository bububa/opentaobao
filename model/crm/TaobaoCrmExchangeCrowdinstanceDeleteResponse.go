package crm

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除人群实例中的指定买家 APIResponse
taobao.crm.exchange.crowdinstance.delete

删除人群实例中的指定买家
*/
type TaobaoCrmExchangeCrowdinstanceDeleteAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoCrmExchangeCrowdinstanceDeleteResponse `json:"crm_exchange_crowdinstance_delete_response,omitempty"` 
    TaobaoCrmExchangeCrowdinstanceDeleteResponse
}

/* model for simplify = false
type TaobaoCrmExchangeCrowdinstanceDeleteResponse struct {

    // 操作成功
    
    SubSuccess   bool `json:"sub_success,omitempty"`
    

}
*/

type TaobaoCrmExchangeCrowdinstanceDeleteResponse struct {

    // 操作成功
    SubSuccess   bool `json:"sub_success,omitempty"`

}
