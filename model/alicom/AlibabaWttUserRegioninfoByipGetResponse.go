package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据ip获取省市信息 APIResponse
alibaba.wtt.user.regioninfo.byip.get

通过ip获取省市信息
*/
type AlibabaWttUserRegioninfoByipGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWttUserRegioninfoByipGetResponse `json:"alibaba_wtt_user_regioninfo_byip_get_response,omitempty"` 
    AlibabaWttUserRegioninfoByipGetResponse
}

/* model for simplify = false
type AlibabaWttUserRegioninfoByipGetResponse struct {

    // 地址信息
    
    Model  *struct {
        RegionInfo  *RegionInfo `json:"region_info,omitempty"`
    } `json:"model,omitempty"`
    

}
*/

type AlibabaWttUserRegioninfoByipGetResponse struct {

    // 地址信息
    Model   *RegionInfo `json:"model,omitempty"`

}
