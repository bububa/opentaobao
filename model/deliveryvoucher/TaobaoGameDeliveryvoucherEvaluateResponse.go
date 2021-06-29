package deliveryvoucher

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
卡券评价回传 APIResponse
taobao.game.deliveryvoucher.evaluate

卡券ISV回传商品评价
*/
type TaobaoGameDeliveryvoucherEvaluateAPIResponse struct {
    model.CommonResponse
    TaobaoGameDeliveryvoucherEvaluateResponse
}

type TaobaoGameDeliveryvoucherEvaluateResponse struct {
    XMLName xml.Name `xml:"game_deliveryvoucher_evaluate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 错误原因
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // code
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 操作状态
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
