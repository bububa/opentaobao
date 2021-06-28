package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘口令配置数据 APIResponse
alibaba.baichuan.taopassword.config

百川淘口令规则配置接口
*/
type AlibabaBaichuanTaopasswordConfigAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_baichuan_taopassword_config_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *ShareResult `json:"result,omitempty" xml:"