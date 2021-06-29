package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据userId(支持单个或批量)获取用户实时位置信息 APIResponse
alibaba.campus.adminmap.userlocationinfo.getactualuserlocationinfobyids

根据userId(支持单个或批量)获取用户实时位置信息 
HSF接口名称：com.alibaba.campus.api.adminmap.service.top.UserLocationQueryApiTopService
HSF方法名称：getActualUserLocationInfoByIds
*/
type AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsResponse
}

type AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_adminmap_userlocationinfo_getactualuserlocationinfobyids_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ListResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
