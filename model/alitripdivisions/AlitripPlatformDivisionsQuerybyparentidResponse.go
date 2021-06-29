package alitripdivisions

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据父节点id查询下级行政区划数据 APIResponse
alitrip.platform.divisions.querybyparentid

根据行政区划id查询下一层级行政区划数据
*/
type AlitripPlatformDivisionsQuerybyparentidAPIResponse struct {
    model.CommonResponse
    AlitripPlatformDivisionsQuerybyparentidResponse
}

type AlitripPlatformDivisionsQuerybyparentidResponse struct {
    XMLName xml.Name `xml:"alitrip_platform_divisions_querybyparentid_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlitripPlatformDivisionsQuerybyparentidResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
