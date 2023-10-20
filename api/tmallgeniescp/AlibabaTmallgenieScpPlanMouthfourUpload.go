package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// Alibabatmallgeniescpplanmouthfourupload 21-M+4PR 回传接口接口
// alibaba.tmallgenie.scp.plan.mouthfour.upload
//
// M+4 PR 回传接口
func Alibabatmallgeniescpplanmouthfourupload(clt *core.SDKClient, req *tmallgeniescp.AlibabatmallgeniescpplanmouthfouruploadAPIRequest, session string) (*tmallgeniescp.AlibabatmallgeniescpplanmouthfouruploadAPIResponse, error) {
	var resp tmallgeniescp.AlibabatmallgeniescpplanmouthfouruploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
