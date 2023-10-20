package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// Alibabatmallgeniescpplansaleqtyget 12-同步销售数据
// alibaba.tmallgenie.scp.plan.sale.qty.get
//
// 同步销售数据
func Alibabatmallgeniescpplansaleqtyget(clt *core.SDKClient, req *tmallgeniescp.AlibabatmallgeniescpplansaleqtygetAPIRequest, session string) (*tmallgeniescp.AlibabatmallgeniescpplansaleqtygetAPIResponse, error) {
	var resp tmallgeniescp.AlibabatmallgeniescpplansaleqtygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
