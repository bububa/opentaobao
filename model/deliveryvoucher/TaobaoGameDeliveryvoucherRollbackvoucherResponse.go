package deliveryvoucher

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
回滚券 APIResponse
taobao.game.deliveryvoucher.rollbackvoucher

提货券发券接口：同步券和订单的关联信息
*/
type TaobaoGameDeliveryvoucherRollbackvoucherAPIResponse struct {
    model.CommonResponse
    TaobaoGameDeliveryvoucherRollbackvoucherResponse
}

type TaobaoGameDeliveryvoucherRollbackvoucherResponse struct {
    XMLName xml.Name `xml:"game_deliveryvoucher_rollbackvoucher_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // message
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // code
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // success
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
