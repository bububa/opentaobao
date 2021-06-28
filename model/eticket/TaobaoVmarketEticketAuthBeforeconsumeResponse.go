package eticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
核销放行的查询接口 APIResponse
taobao.vmarket.eticket.auth.beforeconsume

针对O2O电子凭证核销放行业务，为满足码商能够核销淘宝码而开放的核销查询接口
*/
type TaobaoVmarketEticketAuthBeforeconsumeAPIResponse struct {
    model.CommonResponse
    TaobaoVmarketEticketAuthBeforeconsumeResponse
}

type TaobaoVmarketEticketAuthBeforeconsumeResponse struct {
    XMLName xml.Name `xml:"vmarket_eticket_auth_beforeconsume_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 1:可以进行核销码操作
    
    RetCode   int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`

    
    // 商品标题
    
    ItemTitle   string `json:"item_title,omitempty" xml:"item_title,omitempty"`

    
    // 订单ID
    
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`

    
    // 淘宝卖家ID
    
    TaobaoSid   int64 `json:"taobao_sid,omitempty" xml:"taobao_sid,omitempty"`

    
    // 淘宝卖家旺旺名称
    
    SellerNick   string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`

    
    // 有效期开始时间
    
    ValidStart   string `json:"valid_start,omitempty" xml:"valid_start,omitempty"`

    
    // 有效期结束时间
    
    ValidEnds   string `json:"valid_ends,omitempty" xml:"valid_ends,omitempty"`

    
    // 当前码剩余可核销数量
    
    CodeLeftNum   int64 `json:"code_left_num,omitempty" xml:"code_left_num,omitempty"`

    
}
