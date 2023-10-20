package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// Alibabatmallgeniescpplaninventorqtyget 10-同步库存现有量
// alibaba.tmallgenie.scp.plan.inventor.qty.get
//
// 同步库存现有量
func Alibabatmallgeniescpplaninventorqtyget(clt *core.SDKClient, req *tmallgeniescp.AlibabatmallgeniescpplaninventorqtygetAPIRequest, session string) (*tmallgeniescp.AlibabatmallgeniescpplaninventorqtygetAPIResponse, error) {
	var resp tmallgeniescp.AlibabatmallgeniescpplaninventorqtygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
