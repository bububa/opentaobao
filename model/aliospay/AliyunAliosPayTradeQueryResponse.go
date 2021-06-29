package aliospay

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询支付结果接口 APIResponse
aliyun.alios.pay.trade.query

商户用来查询支付结果接口
*/
type AliyunAliosPayTradeQueryAPIResponse struct {
    model.CommonResponse
    AliyunAliosPayTradeQueryResponse
}

type AliyunAliosPayTradeQueryResponse struct {
    XMLName xml.Name `xml:"aliyun_alios_pay_trade_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 响应参数
    
    AliospayResponse   *AliOSPayResponse `json:"aliospay_response,omitempty" xml:"aliospay_response,omitempty"`

    
}
