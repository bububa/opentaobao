package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
交易迎客松通用服务接口 APIResponse
yunos.trade.admin.common.operation

迎客松交易相关通用接口
*/
type YunosTradeAdminCommonOperationAPIResponse struct {
    model.CommonResponse
    YunosTradeAdminCommonOperationResponse
}

type YunosTradeAdminCommonOperationResponse struct {
    XMLName xml.Name `xml:"yunos_trade_admin_common_operation_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TopResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
