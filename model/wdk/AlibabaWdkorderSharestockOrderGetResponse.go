package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
猫超商户订单拉取 APIResponse
alibaba.wdkorder.sharestock.order.get

商户拉取猫超订单数据
*/
type AlibabaWdkorderSharestockOrderGetAPIResponse struct {
    model.CommonResponse
    AlibabaWdkorderSharestockOrderGetResponse
}

type AlibabaWdkorderSharestockOrderGetResponse struct {
    XMLName xml.Name `xml:"alibaba_wdkorder_sharestock_order_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 调用结果
    
    Result   *MaochaoOrderQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
