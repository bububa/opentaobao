package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商户回传餐饮加工单状态 APIResponse
alibaba.wdkopen.cateorder.pull

商户回传餐饮加工单状态
*/
type AlibabaWdkopenCateorderPullAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkopenCateorderPullResponse `json:"alibaba_wdkopen_cateorder_pull_response,omitempty"` 
    AlibabaWdkopenCateorderPullResponse
}

/* model for simplify = false
type AlibabaWdkopenCateorderPullResponse struct {

    // 调用返回
    
    TopBaseResult  *struct {
        TopBaseResult  *TopBaseResult `json:"top_base_result,omitempty"`
    } `json:"top_base_result,omitempty"`
    

}
*/

type AlibabaWdkopenCateorderPullResponse struct {

    // 调用返回
    TopBaseResult   *TopBaseResult `json:"top_base_result,omitempty"`

}
