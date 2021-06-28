package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
渠道间库存规则设置接口 APIResponse
taobao.qimen.inventoryrule.create

渠道间库存规则设置
*/
type TaobaoQimenInventoryruleCreateAPIResponse struct {
    model.CommonResponse
    TaobaoQimenInventoryruleCreateResponse
}

type TaobaoQimenInventoryruleCreateResponse struct {
    XMLName xml.Name `xml:"qimen_inventoryrule_create_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *ResponseDO `json:"response,omitempty" xml:"response,omitempty"`

    
}
