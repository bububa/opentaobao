package opentrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
组团购场景参团 APIResponse
taobao.opentrade.group.join

组团购场景下，用户参团
*/
type TaobaoOpentradeGroupJoinAPIResponse struct {
    model.CommonResponse
    TaobaoOpentradeGroupJoinResponse
}

type TaobaoOpentradeGroupJoinResponse struct {
    XMLName xml.Name `xml:"opentrade_group_join_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

}
