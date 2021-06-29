package crm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
向活动人群实例中增加买家 API返回值 
taobao.crm.exchange.crowdinstance.add

向活动人群实例中增加买家
*/
type TaobaoCrmExchangeCrowdinstanceAddAPIResponse struct {
    model.CommonResponse
    TaobaoCrmExchangeCrowdinstanceAddResponse
}

// 向活动人群实例中增加买家 成功返回结果
type TaobaoCrmExchangeCrowdinstanceAddResponse struct {
    XMLName xml.Name `xml:"crm_exchange_crowdinstance_add_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 调用是否成功
    SubSuccess   bool `json:"sub_success,omitempty" xml:"sub_success,omitempty"`
}
