package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosArkHealthOfflineLevelGetAPIRequest
获取mall的离线等级 API请求
alibaba.mos.ark.health.offline.level.get

获取mall的离线等级 */
type AlibabaMosArkHealthOfflineLevelGetAPIRequest struct {
	model.Params
	// 商场id
	_mallId string
}

// New
