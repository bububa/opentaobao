package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据分组条件查询分组下的空间单元不包涵业务属性信息 API返回值 
alibaba.campus.adminmap.poiinfo.getlistbygroup

根据分组条件查询分组下的空间单元不包涵业务属性信息
*/
type AlibabaCampusAdminmapPoiinfoGetlistbygroupAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAdminmapPoiinfoGetlistbygroupResponse
}

// 根据分组条件查询分组下的空间单元不包涵业务属性信息 成功返回结果
type AlibabaCampusAdminmapPoiinfoGetlistbygroupResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_adminmap_poiinfo_getlistbygroup_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}
