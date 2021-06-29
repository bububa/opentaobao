package cainiaohandover

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取消交接单 APIResponse
cainiao.global.handover.cancel

提供给ISV通过该接口取消交接单
*/
type CainiaoGlobalHandoverCancelAPIResponse struct {
    model.CommonResponse
    CainiaoGlobalHandoverCancelResponse
}

type CainiaoGlobalHandoverCancelResponse struct {
    XMLName xml.Name `xml:"cainiao_global_handover_cancel_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 请求结果
    
    Result   *HsfResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
