package crm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
向活动人群实例中增加买家 APIResponse
taobao.crm.exchange.crowdinstance.add

向活动人群实例中增加买家
*/
type TaobaoCrmExchangeCrowdinstanceAddAPIResponse struct {
    model.CommonResponse
    TaobaoCrmExchangeCrowdinstanceAddResponse
}

type TaobaoCrmExchangeCrowdinstanceAddResponse struct {
    XMLName xml.Name `xml:"crm_exchange_crowdinstance_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 调用是否成功
    
    SubSuccess   bool `json:"sub_success,omitempty" xml:"sub_success,omitempty"`

    
}
