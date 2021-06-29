package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】查询审批单 APIResponse
alitrip.btrip.corpop.apply.get

【商旅】查询审批单
*/
type AlitripBtripCorpopApplyGetAPIResponse struct {
    model.CommonResponse
    AlitripBtripCorpopApplyGetResponse
}

type AlitripBtripCorpopApplyGetResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_corpop_apply_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 出参对象
    
    Result   *HisvResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
