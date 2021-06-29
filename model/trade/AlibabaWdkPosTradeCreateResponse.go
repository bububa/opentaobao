package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
轻pos品牌营销下单接口 APIResponse
alibaba.wdk.pos.trade.create

提供给石基进行轻pos品牌营销下单
*/
type AlibabaWdkPosTradeCreateAPIResponse struct {
    model.CommonResponse
    AlibabaWdkPosTradeCreateResponse
}

type AlibabaWdkPosTradeCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_pos_trade_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 创单结果
    
    Result   *FastBuyPosCreateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
