package crm

import (
    "github.com/bububa/opentaobao/model"
)

/* 
向活动人群实例中增加买家 APIResponse
taobao.crm.exchange.crowdinstance.add

向活动人群实例中增加买家
*/
type TaobaoCrmExchangeCrowdinstanceAddAPIResponse struct {
    model.CommonResponse
    Response *TaobaoCrmExchangeCrowdinstanceAddResponse `json:"taobao_crm_exchange_crowdinstance_add_response,omitempty"`
}

type TaobaoCrmExchangeCrowdinstanceAddResponse struct {

    // 调用是否成功
    SubSuccess   bool `json:"sub_success,omitempty"`

}
