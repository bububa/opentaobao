package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
五道口库存发布接口（针对外部渠道） APIResponse
alibaba.wdk.stock.publish

五道口库存发布接口（针对外部渠道）
*/
type AlibabaWdkStockPublishAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkStockPublishResponse `json:"alibaba_wdk_stock_publish_response,omitempty"`
}

type AlibabaWdkStockPublishResponse struct {

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

    // errorCode
    MsgCode   string `json:"msg_code,omitempty"`

    // errorMsg
    Message   string `json:"message,omitempty"`

}
