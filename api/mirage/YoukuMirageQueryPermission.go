package mirage

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mirage"
)

// YoukuMirageQueryPermission 优酷播控查询是否可播API
// youku.mirage.query.permission
//
// 根据节目ID或者VID查询视频或者节目是否可以播放
func YoukuMirageQueryPermission(ctx context.Context, clt *core.SDKClient, req *mirage.YoukuMirageQueryPermissionAPIRequest, resp *mirage.YoukuMirageQueryPermissionAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
