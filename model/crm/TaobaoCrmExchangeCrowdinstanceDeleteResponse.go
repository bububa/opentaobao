package crm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除人群实例中的指定买家 APIResponse
taobao.crm.exchange.crowdinstance.delete

删除人群实例中的指定买家
*/
type TaobaoCrmExchangeCrowdinstanceDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoCrmExchangeCrowdinstanceDeleteResponse
}

type TaobaoCrmExchangeCrowdinstanceDeleteResponse struct {
    XMLName xml.Name `xml:"crm_exchange_crowdinstance_delete_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 操作成功
    
    SubSuccess   bool `json:"sub_success,omitempty" xml:"sub_success,omitempty"`

    
}
