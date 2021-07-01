package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

/* AlibabaMosArkHealthOfflineLevelGet
获取mall的离线等级
alibaba.mos.ark.health.offline.level.get

获取mall的离线等级 */
func AlibabaMosArkHealthOfflineLevelGet(clt *core.SDKClient, req *mos.AlibabaMosArkHealthOfflineLevelGetAPIRequest, session string) (*mos.AlibabaMosArkHealthOfflineLevelGetAPIResponse, error) {
	var resp mos.AlibabaMosArkHealthOfflineLevelGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
