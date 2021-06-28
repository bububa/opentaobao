package drug

import (
    "github.com/bububa/opentaobao/model"
)

/* 
o2o查看处方图片 APIResponse
alibaba.alihealth.nr.rx.queryimage

o2o商家查看处方图片，包括电子图片与纸质图片
*/
type AlibabaAlihealthNrRxQueryimageAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlihealthNrRxQueryimageResponse `json:"alibaba_alihealth_nr_rx_queryimage_response,omitempty"` 
    AlibabaAlihealthNrRxQueryimageResponse
}

/* model for simplify = false
type AlibabaAlihealthNrRxQueryimageResponse struct {

    // 成功或失败
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 错误编码
    
    ErrorInfoCode   string `json:"error_info_code,omitempty"`
    

    // 错误描述
    
    ErrorInfoMsg   string `json:"error_info_msg,omitempty"`
    

    // 处方图片url，多张图片用逗号分隔
    
    Result   string `json:"result,omitempty"`
    

}
*/

type AlibabaAlihealthNrRxQueryimageResponse struct {

    // 成功或失败
    IsSuccess   bool `json:"is_success,omitempty"`

    // 错误编码
    ErrorInfoCode   string `json:"error_info_code,omitempty"`

    // 错误描述
    ErrorInfoMsg   string `json:"error_info_msg,omitempty"`

    // 处方图片url，多张图片用逗号分隔
    Result   string `json:"result,omitempty"`

}
