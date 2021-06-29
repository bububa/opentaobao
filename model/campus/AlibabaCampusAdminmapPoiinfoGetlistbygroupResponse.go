package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据分组条件查询分组下的空间单元不包涵业务属性信息 APIResponse
alibaba.campus.adminmap.poiinfo.getlistbygroup

根据分组条件查询分组下的空间单元不包涵业务属性信息
*/
type AlibabaCampusAdminmapPoiinfoGetlistbygroupAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAdminmapPoiinfoGetlistbygroupResponse
}

type AlibabaCampusAdminmapPoiinfoGetlistbygroupResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_adminmap_poiinfo_getlistbygroup_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ListResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
