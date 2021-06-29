package mirage

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
优酷播控查询是否可播API APIResponse
youku.mirage.query.permission

根据节目ID或者VID查询视频或者节目是否可以播放
*/
type YoukuMirageQueryPermissionAPIResponse struct {
    model.CommonResponse
    YoukuMirageQueryPermissionResponse
}

type YoukuMirageQueryPermissionResponse struct {
    XMLName xml.Name `xml:"youku_mirage_query_permission_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回值
    
    Resp   *PermissionResponseDto `json:"resp,omitempty" xml:"resp,omitempty"`

    
}
