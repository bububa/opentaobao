package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// Alibabatmallgeniescpplanrawpogapreturn 二级物料-LT内的POGAP数据回传
// alibaba.tmallgenie.scp.plan.rawpo.gap.return
//
// 二级物料-LT内的POGAP数据回传
func Alibabatmallgeniescpplanrawpogapreturn(clt *core.SDKClient, req *tmallgeniescp.AlibabatmallgeniescpplanrawpogapreturnAPIRequest, session string) (*tmallgeniescp.AlibabatmallgeniescpplanrawpogapreturnAPIResponse, error) {
	var resp tmallgeniescp.AlibabatmallgeniescpplanrawpogapreturnAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
