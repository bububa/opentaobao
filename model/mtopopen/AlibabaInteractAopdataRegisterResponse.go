package mtopopen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
资源位数据推送接口 APIResponse
alibaba.interact.aopdata.register

提供给isv，查询以及推送浮层资源位的三方活动数据
*/
type AlibabaInteractAopdataRegisterAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaInteractAopdataRegisterResponse `json:"alibaba_interact_aopdata_register_response,omitempty"` 
    AlibabaInteractAopdataRegisterResponse
}

/* model for simplify = false
type AlibabaInteractAopdataRegisterResponse struct {

    // 接口返回model
    
    Result  *struct {
        AlibabaInteractAopdataRegisterResult  *AlibabaInteractAopdataRegisterResult `json:"alibaba_interact_aopdata_register_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaInteractAopdataRegisterResponse struct {

    // 接口返回model
    Result   *AlibabaInteractAopdataRegisterResult `json:"result,omitempty"`

}
