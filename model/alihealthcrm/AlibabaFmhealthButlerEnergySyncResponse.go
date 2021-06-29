package alihealthcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
同步用户消耗能量 APIResponse
alibaba.fmhealth.butler.energy.sync

同步用户消耗能量，用户消耗s点或卡路里后，同步给健康平台
*/
type AlibabaFmhealthButlerEnergySyncAPIResponse struct {
    model.CommonResponse
    AlibabaFmhealthButlerEnergySyncResponse
}

type AlibabaFmhealthButlerEnergySyncResponse struct {
    XMLName xml.Name `xml:"alibaba_fmhealth_butler_energy_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
