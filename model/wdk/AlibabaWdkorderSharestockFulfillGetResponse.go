package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商户订单履约数据获取 APIResponse
alibaba.wdkorder.sharestock.fulfill.get

商户订单履约数据获取
*/
type AlibabaWdkorderSharestockFulfillGetAPIResponse struct {
    model.CommonResponse
    AlibabaWdkorderSharestockFulfillGetResponse
}

type AlibabaWdkorderSharestockFulfillGetResponse struct {
    XMLName xml.Name `xml:"alibaba_wdkorder_sharestock_fulfill_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 调用结果
    
    Result   *MaochaoOrderFulfillQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
