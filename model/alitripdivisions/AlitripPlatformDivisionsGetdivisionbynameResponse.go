package alitripdivisions

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据中文名称与行政区划级别查询行政区划数据 API返回值 
alitrip.platform.divisions.getdivisionbyname

根据中文名称与行政区划级别查询行政区划数据
*/
type AlitripPlatformDivisionsGetdivisionbynameAPIResponse struct {
    model.CommonResponse
    AlitripPlatformDivisionsGetdivisionbynameResponse
}

// 根据中文名称与行政区划级别查询行政区划数据 成功返回结果
type AlitripPlatformDivisionsGetdivisionbynameResponse struct {
    XMLName xml.Name `xml:"alitrip_platform_divisions_getdivisionbyname_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Results   []string `json:"results,omitempty" xml:"results>string,omitempty"`
}
