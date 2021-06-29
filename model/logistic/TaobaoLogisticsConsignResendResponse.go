package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改物流公司和运单号 APIResponse
taobao.logistics.consign.resend

支持卖家发货后修改运单号;支持在线下单和自己联系两种发货方式;使用条件：<br><br/>1、必须是已发货订单，自己联系发货的必须24小时内才可修改；在线下单的，必须下单后物流公司未揽收成功前才可修改；<br/>2、自己联系只能切换为自己联系的公司，在线下单也只能切换为在线下单的物流公司。
*/
type TaobaoLogisticsConsignResendAPIResponse struct {
    model.CommonResponse
    TaobaoLogisticsConsignResendResponse
}

type TaobaoLogisticsConsignResendResponse struct {
    XMLName xml.Name `xml:"logistics_consign_resend_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回发货是否成功is_success
    
    Shipping   *Shipping `json:"shipping,omitempty" xml:"shipping,omitempty"`

    
}
