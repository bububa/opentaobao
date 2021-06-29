package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
上传用户实时位置 APIResponse
alibaba.campus.adminmap.userlocationinfo.insertactualuserlocationinfo

上传用户实时位置
HSF接口名称：com.alibaba.campus.api.adminmap.service.top.UserLocationQueryApiTopService
HSF方法名称：insertActualUserLocationInfo
*/
type AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoResponse
}

type AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_adminmap_userlocationinfo_insertactualuserlocationinfo_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
