package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
导购员判断 APIResponse
taobao.omniorder.print.sale.judge

用于判断当前子账号是否导购员
*/
type TaobaoOmniorderPrintSaleJudgeAPIResponse struct {
    model.CommonResponse
    TaobaoOmniorderPrintSaleJudgeResponse
}

type TaobaoOmniorderPrintSaleJudgeResponse struct {
    XMLName xml.Name `xml:"omniorder_print_sale_judge_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // data
    
    Data   string `json:"data,omitempty" xml:"data,omitempty"`

    
}
