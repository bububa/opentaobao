package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】修改出差审批单（行程） APIResponse
alitrip.btrip.corpop.apply.modify

【商旅】修改出差审批单（行程）
*/
type AlitripBtripCorpopApplyModifyAPIResponse struct {
    model.CommonResponse
    AlitripBtripCorpopApplyModifyResponse
}

type AlitripBtripCorpopApplyModifyResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_corpop_apply_modify_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果对象
    
    Result   *HisvResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
