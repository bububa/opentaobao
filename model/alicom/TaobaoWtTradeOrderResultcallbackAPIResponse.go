package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家回调接口 API返回值 
taobao.wt.trade.order.resultcallback

阿里通信定制服务，商家发货后进行调用该接口，用于自动发货并确认收货
*/
type TaobaoWtTradeOrderResultcallbackAPIResponse struct {
    model.CommonResponse
    TaobaoWtTradeOrderResultcallbackAPIResponseModel
}

// 商家回调接口 成功返回结果
type TaobaoWtTradeOrderResultcallbackAPIResponseModel struct {
    XMLName xml.Name `xml:"wt_trade_order_resultcallback_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *CommonRtnDo `json:"result,omitempty" xml:"result,omitempty"`
}
