package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索审批单 APIResponse
alitrip.btrip.apply.search

外部企业调用获取本企业审批单列表数据
*/
type AlitripBtripApplySearchAPIResponse struct {
    model.CommonResponse
    AlitripBtripApplySearchResponse
}

type AlitripBtripApplySearchResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_apply_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *BtriphomeResultSupport `json:"result,omitempty" xml:"result,omitempty"`

    
}
