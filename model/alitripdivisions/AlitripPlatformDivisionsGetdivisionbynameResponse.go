package alitripdivisions

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据中文名称与行政区划级别查询行政区划数据 APIResponse
alitrip.platform.divisions.getdivisionbyname

根据中文名称与行政区划级别查询行政区划数据
*/
type AlitripPlatformDivisionsGetdivisionbynameAPIResponse struct {
    model.CommonResponse
    AlitripPlatformDivisionsGetdivisionbynameResponse
}

type AlitripPlatformDivisionsGetdivisionbynameResponse struct {
    XMLName xml.Name `xml:"alitrip_platform_divisions_getdivisionbyname_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Results   []string `json:"results,omitempty" xml:"results>string,omitempty"`
    
    
}
