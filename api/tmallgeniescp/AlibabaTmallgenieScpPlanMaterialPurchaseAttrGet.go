package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// Alibabatmallgeniescpplanmaterialpurchaseattrget 物料的采购属性查询
// alibaba.tmallgenie.scp.plan.material.purchase.attr.get
//
// 物料的采购属性查询
func Alibabatmallgeniescpplanmaterialpurchaseattrget(clt *core.SDKClient, req *tmallgeniescp.AlibabatmallgeniescpplanmaterialpurchaseattrgetAPIRequest, session string) (*tmallgeniescp.AlibabatmallgeniescpplanmaterialpurchaseattrgetAPIResponse, error) {
	var resp tmallgeniescp.AlibabatmallgeniescpplanmaterialpurchaseattrgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
