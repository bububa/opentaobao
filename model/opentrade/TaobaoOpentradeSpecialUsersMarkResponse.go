package opentrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
专属下单可购买用户标记 APIResponse
taobao.opentrade.special.users.mark

对于专属下单的交易场景，根据openid标记用户可购买商品
*/
type TaobaoOpentradeSpecialUsersMarkAPIResponse struct {
    model.CommonResponse
    TaobaoOpentradeSpecialUsersMarkResponse
}

type TaobaoOpentradeSpecialUsersMarkResponse struct {
    XMLName xml.Name `xml:"opentrade_special_users_mark_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 标记成功的用户数
    
    Result   int64 `json:"result,omitempty" xml:"result,omitempty"`

    
}
