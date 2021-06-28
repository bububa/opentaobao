package baichuan

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘口令配置数据 APIResponse
alibaba.baichuan.taopassword.config

百川淘口令规则配置接口
*/
type AlibabaBaichuanTaopasswordConfigAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaBaichuanTaopasswordConfigResponse `json:"alibaba_baichuan_taopassword_config_response,omitempty"` 
    AlibabaBaichuanTaopasswordConfigResponse
}

/* model for simplify = false
type AlibabaBaichuanTaopasswordConfigResponse struct {

    // result
    
    Result  *struct {
        ShareResult  *ShareResult `json:"share_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaBaichuanTaopasswordConfigResponse struct {

    // result
    Result   *ShareResult `json:"result,omitempty"`

}
