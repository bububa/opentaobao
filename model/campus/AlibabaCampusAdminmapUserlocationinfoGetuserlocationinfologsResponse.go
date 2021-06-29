package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分时间段获取用户历史位置信息 APIResponse
alibaba.campus.adminmap.userlocationinfo.getuserlocationinfologs

分时间段获取用户历史位置信息
*/
type AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsResponse
}

type AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_adminmap_userlocationinfo_getuserlocationinfologs_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ListResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
