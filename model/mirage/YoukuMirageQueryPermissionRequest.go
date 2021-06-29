package mirage

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优酷播控查询是否可播API APIRequest
youku.mirage.query.permission

根据节目ID或者VID查询视频或者节目是否可以播放
*/
type YoukuMirageQueryPermissionRequest struct {
    model.Params

    // 入参
    permissionRequestDto   *PermissionRequestDto 

}

func NewYoukuMirageQueryPermissionRequest() *YoukuMirageQueryPermissionRequest{
    return &YoukuMirageQueryPermissionRequest{
        Params: model.NewParams(),
    }
}

func (r YoukuMirageQueryPermissionRequest) GetApiMethodName() string {
    return "youku.mirage.query.permission"
}

func (r YoukuMirageQueryPermissionRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YoukuMirageQueryPermissionRequest) SetPermissionRequestDto(permissionRequestDto *PermissionRequestDto) error {
    r.permissionRequestDto = permissionRequestDto
    r.Set("permission_request_dto", permissionRequestDto)
    return nil
}

func (r YoukuMirageQueryPermissionRequest) GetPermissionRequestDto() *PermissionRequestDto {
    return r.permissionRequestDto
}

