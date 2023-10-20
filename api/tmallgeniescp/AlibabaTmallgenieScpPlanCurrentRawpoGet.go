package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// Alibabatmallgeniescpplancurrentrawpoget 二级物料-PO数据同步
// alibaba.tmallgenie.scp.plan.current.rawpo.get
//
// 二级物料-PO数据同步（WO-W[TL])
func Alibabatmallgeniescpplancurrentrawpoget(clt *core.SDKClient, req *tmallgeniescp.AlibabatmallgeniescpplancurrentrawpogetAPIRequest, session string) (*tmallgeniescp.AlibabatmallgeniescpplancurrentrawpogetAPIResponse, error) {
	var resp tmallgeniescp.AlibabatmallgeniescpplancurrentrawpogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
