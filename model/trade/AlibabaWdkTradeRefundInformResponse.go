package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
外部渠道通知淘鲜达退款成功接口 APIResponse
alibaba.wdk.trade.refund.inform

该接口用于外部渠道退款成功后，通知淘鲜达底层履约完成退款流程。
*/
type AlibabaWdkTradeRefundInformAPIResponse struct {
    model.CommonResponse
    AlibabaWdkTradeRefundInformResponse
}

type AlibabaWdkTradeRefundInformResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_trade_refund_inform_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *InformRefundSuccessResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
