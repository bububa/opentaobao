package pur

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// Alibabaceressupplierpoquerydetail 采购供应商订单明细查询接口
// alibaba.ceres.supplier.po.querydetail
//
// 采购供应商订单明细查询接口
func Alibabaceressupplierpoquerydetail(clt *core.SDKClient, req *pur.AlibabaceressupplierpoquerydetailAPIRequest, session string) (*pur.AlibabaceressupplierpoquerydetailAPIResponse, error) {
	var resp pur.AlibabaceressupplierpoquerydetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
