package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// Alibabatmallgeniescpplansaleforcastsalerupload 19-销售预测数量（销管）回传接口
// alibaba.tmallgenie.scp.plan.saleforcast.saler.upload
//
// 销售预测数量（销管）回传接口
func Alibabatmallgeniescpplansaleforcastsalerupload(clt *core.SDKClient, req *tmallgeniescp.AlibabatmallgeniescpplansaleforcastsaleruploadAPIRequest, session string) (*tmallgeniescp.AlibabatmallgeniescpplansaleforcastsaleruploadAPIResponse, error) {
	var resp tmallgeniescp.AlibabatmallgeniescpplansaleforcastsaleruploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
