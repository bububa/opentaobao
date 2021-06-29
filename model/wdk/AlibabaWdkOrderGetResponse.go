package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
交易订单详情查询 APIResponse
alibaba.wdk.order.get

五道口三江单据查询接口
*/
type AlibabaWdkOrderGetAPIResponse struct {
    model.CommonResponse
    AlibabaWdkOrderGetResponse
}

type AlibabaWdkOrderGetResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_order_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回数据
    
    Result   *AlibabaWdkOrderGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
