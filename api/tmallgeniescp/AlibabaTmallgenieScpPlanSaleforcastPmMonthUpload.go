package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// Alibabatmallgeniescpplansaleforcastpmmonthupload 24-销售月预测数量（产管）回传-月度
// alibaba.tmallgenie.scp.plan.saleforcast.pm.month.upload
//
// 销售月预测数量（产管）回传-月度
func Alibabatmallgeniescpplansaleforcastpmmonthupload(clt *core.SDKClient, req *tmallgeniescp.AlibabatmallgeniescpplansaleforcastpmmonthuploadAPIRequest, session string) (*tmallgeniescp.AlibabatmallgeniescpplansaleforcastpmmonthuploadAPIResponse, error) {
	var resp tmallgeniescp.AlibabatmallgeniescpplansaleforcastpmmonthuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
