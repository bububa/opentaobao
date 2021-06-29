package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】isv添加审批单 APIResponse
alitrip.btrip.corpop.apply.add

【商旅】isv添加审批单
*/
type AlitripBtripCorpopApplyAddAPIResponse struct {
    model.CommonResponse
    AlitripBtripCorpopApplyAddResponse
}

type AlitripBtripCorpopApplyAddResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_corpop_apply_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 出参数
    
    Result   *HisvResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
