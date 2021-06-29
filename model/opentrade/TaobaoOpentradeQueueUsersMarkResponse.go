package opentrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
尖货交易可购买用户标记 APIResponse
taobao.opentrade.queue.users.mark

尖货交易用户标记信息回传，根据openId标记用户可购买商品
*/
type TaobaoOpentradeQueueUsersMarkAPIResponse struct {
    model.CommonResponse
    TaobaoOpentradeQueueUsersMarkResponse
}

type TaobaoOpentradeQueueUsersMarkResponse struct {
    XMLName xml.Name `xml:"opentrade_queue_users_mark_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 标记成功的用户数
    
    Result   int64 `json:"result,omitempty" xml:"result,omitempty"`

    
}
