package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
确认发货通知接口 APIResponse
taobao.logistics.online.confirm

<br><font color='red'>仅在使用taobao.logistics.online.send 发货时未输入运单号的情况下，需要使用该接口补充填写运单号，来确认发货。<br>
确认发货的目的是让交易流程继续走下去，确认发货后交易状态会由【买家已付款】变为【卖家已发货】。</font>
*/
type TaobaoLogisticsOnlineConfirmAPIResponse struct {
    model.CommonResponse
    TaobaoLogisticsOnlineConfirmResponse
}

type TaobaoLogisticsOnlineConfirmResponse struct {
    XMLName xml.Name `xml:"logistics_online_confirm_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 只返回is_success：是否成功。
    
    Shipping   *Shipping `json:"shipping,omitempty" xml:"shipping,omitempty"`

    
}
