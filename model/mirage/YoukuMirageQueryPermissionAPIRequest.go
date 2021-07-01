package mirage

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YoukuMirageQueryPermissionAPIRequest
优酷播控查询是否可播API API请求
youku.mirage.query.permission

根据节目ID或者VID查询视频或者节目是否可以播放 */
type YoukuMirageQueryPermissionAPIRequest struct {
	model.Params
	// 入参
	_permissionRequestDto *PermissionRequestDto
}

// New
