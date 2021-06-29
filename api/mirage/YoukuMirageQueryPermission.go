package mirage

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mirage"
)

/* 
优酷播控查询是否可播API 
youku.mirage.query.permission

根据节目ID或者VID查询视频或者节目是否可以播放
*/
func YoukuMirageQueryPermission(clt *core.SDKClient, req *mirage.YoukuMirageQueryPermissionRequest, session string) (*mirage.YoukuMirageQueryPermissionAPIResponse, error) {
    var resp mirage.YoukuMirageQueryPermissionAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
