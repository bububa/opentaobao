package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// Alibabatmallgeniescpplanmaterielget 1-IBP同步物料接口
// alibaba.tmallgenie.scp.plan.materiel.get
//
// IBP同步物料接口
func Alibabatmallgeniescpplanmaterielget(clt *core.SDKClient, req *tmallgeniescp.AlibabatmallgeniescpplanmaterielgetAPIRequest, session string) (*tmallgeniescp.AlibabatmallgeniescpplanmaterielgetAPIResponse, error) {
	var resp tmallgeniescp.AlibabatmallgeniescpplanmaterielgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
