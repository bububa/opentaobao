package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口库存发布接口（针对外部渠道） APIResponse
alibaba.wdk.stock.publish

五道口库存发布接口（针对外部渠道）
*/
type AlibabaWdkStockPublishAPIResponse struct {
    model.CommonResponse
    AlibabaWdkStockPublishResponse
}

type AlibabaWdkStockPublishResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_stock_publish_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // success
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // errorCode
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
    // errorMsg
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
}
