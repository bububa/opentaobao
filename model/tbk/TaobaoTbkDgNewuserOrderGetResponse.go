package tbk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-新用户订单明细查询 APIResponse
taobao.tbk.dg.newuser.order.get

拉新API
*/
type TaobaoTbkDgNewuserOrderGetAPIResponse struct {
    model.CommonResponse
    TaobaoTbkDgNewuserOrderGetResponse
}

type TaobaoTbkDgNewuserOrderGetResponse struct {
    XMLName xml.Name `xml:"tbk_dg_newuser_order_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // data
    
    Results   *TaobaoTbkDgNewuserOrderGetResults `json:"results,omitempty" xml:"results,omitempty"`

    
}
