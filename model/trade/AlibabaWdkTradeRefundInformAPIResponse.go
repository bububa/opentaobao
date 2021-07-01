package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
外部渠道通知淘鲜达退款成功接口 API返回值 
alibaba.wdk.trade.refund.inform

该接口用于外部渠道退款成功后，通知淘鲜达底层履约完成退款流程。
*/
type AlibabaWdkTradeRefundInformAPIResponse struct {
    model.CommonResponse
    AlibabaWdkTradeRefundInformAPIResponseModel
}

// 外部渠道通知淘鲜达退款成功接口 成功返回结果
type AlibabaWdkTradeRefundInformAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_wdk_trade_refund_inform_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *InformRefundSuccessResult `json:"result,omitempty" xml:"result,omitempty"`
}
