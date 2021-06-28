package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商户回传餐饮加工单状态 APIResponse
alibaba.wdkopen.cateorder.pull

商户回传餐饮加工单状态
*/
type AlibabaWdkopenCateorderPullAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdkopen_cateorder_pull_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用返回
    
    TopBaseResult   *TopBaseResult `json:"top_base_result,omitempty" xml:"