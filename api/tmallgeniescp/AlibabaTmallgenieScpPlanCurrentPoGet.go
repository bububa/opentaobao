package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// Alibabatmallgeniescpplancurrentpoget 11-同步本周的po单（从W-1周到W+4周）
// alibaba.tmallgenie.scp.plan.current.po.get
//
// 11-同步本周的po单（从W-1周到W+4周）
func Alibabatmallgeniescpplancurrentpoget(clt *core.SDKClient, req *tmallgeniescp.AlibabatmallgeniescpplancurrentpogetAPIRequest, session string) (*tmallgeniescp.AlibabatmallgeniescpplancurrentpogetAPIResponse, error) {
	var resp tmallgeniescp.AlibabatmallgeniescpplancurrentpogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
