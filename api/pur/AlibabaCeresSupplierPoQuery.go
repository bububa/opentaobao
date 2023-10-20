package pur

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// Alibabaceressupplierpoquery 采购供应商订单查询接口
// alibaba.ceres.supplier.po.query
//
// 采购供应商订单查询接口
func Alibabaceressupplierpoquery(clt *core.SDKClient, req *pur.AlibabaceressupplierpoqueryAPIRequest, session string) (*pur.AlibabaceressupplierpoqueryAPIResponse, error) {
	var resp pur.AlibabaceressupplierpoqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
