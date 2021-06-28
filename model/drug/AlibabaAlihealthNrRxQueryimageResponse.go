package drug

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
o2o查看处方图片 APIResponse
alibaba.alihealth.nr.rx.queryimage

o2o商家查看处方图片，包括电子图片与纸质图片
*/
type AlibabaAlihealthNrRxQueryimageAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_alihealth_nr_rx_queryimage_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 成功或失败
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"