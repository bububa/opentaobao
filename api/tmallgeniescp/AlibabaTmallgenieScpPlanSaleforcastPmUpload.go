package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// Alibabatmallgeniescpplansaleforcastpmupload 18-销售预测数量（产管）回传接口
// alibaba.tmallgenie.scp.plan.saleforcast.pm.upload
//
// 销售预测数量（产管）回传接口
func Alibabatmallgeniescpplansaleforcastpmupload(clt *core.SDKClient, req *tmallgeniescp.AlibabatmallgeniescpplansaleforcastpmuploadAPIRequest, session string) (*tmallgeniescp.AlibabatmallgeniescpplansaleforcastpmuploadAPIResponse, error) {
	var resp tmallgeniescp.AlibabatmallgeniescpplansaleforcastpmuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
