package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单个审批单 APIResponse
alitrip.btrip.apply.get

获取单个审批单的详情数据
*/
type AlitripBtripApplyGetAPIResponse struct {
    model.CommonResponse
    AlitripBtripApplyGetResponse
}

type AlitripBtripApplyGetResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_apply_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *BtriphomeResultSupport `json:"result,omitempty" xml:"result,omitempty"`

    
}
