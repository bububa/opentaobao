package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
退款信息审核 APIResponse
taobao.rdc.aligenius.refunds.check

根据退款信息，对退款单进行审核
*/
type TaobaoRdcAligeniusRefundsCheckAPIResponse struct {
    model.CommonResponse
    TaobaoRdcAligeniusRefundsCheckResponse
}

type TaobaoRdcAligeniusRefundsCheckResponse struct {
    XMLName xml.Name `xml:"rdc_aligenius_refunds_check_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TaobaoRdcAligeniusRefundsCheckResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
