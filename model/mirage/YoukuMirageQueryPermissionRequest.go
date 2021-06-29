package mirage

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优酷播控查询是否可播API API请求
youku.mirage.query.permission

根据节目ID或者VID查询视频或者节目是否可以播放
*/
type YoukuMirageQueryPermissionRequest struct {
    model.Params
    // 入参
    _permissionRequestDto   *PermissionRequestDto
}

// 初始化YoukuMirageQueryPermissionRequest对象
func NewYoukuMirageQueryPermissionRequest() *YoukuMirageQueryPermissionRequest{
    return &YoukuMirageQueryPermissionRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YoukuMirageQueryPermissionRequest) GetApiMethodName() string {
    return "youku.mirage.query.permission"
}

// IRequest interface 方法, 获取API参数
func (r YoukuMirageQueryPermissionRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PermissionRequestDto Setter
// 入参
func (r *YoukuMirageQueryPermissionRequest) SetPermissionRequestDto(_permissionRequestDto *PermissionRequestDto) error {
    r._permissionRequestDto = _permissionRequestDto
    r.Set("permission_request_dto", _permissionRequestDto)
    return nil
}

// PermissionRequestDto Getter
func (r YoukuMirageQueryPermissionRequest) GetPermissionRequestDto() *PermissionRequestDto {
    return r._permissionRequestDto
}
