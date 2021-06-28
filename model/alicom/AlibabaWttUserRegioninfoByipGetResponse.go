package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据ip获取省市信息 APIResponse
alibaba.wtt.user.regioninfo.byip.get

通过ip获取省市信息
*/
type AlibabaWttUserRegioninfoByipGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wtt_user_regioninfo_byip_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 地址信息
    
    Model   *RegionInfo `json:"model,omitempty" xml:"