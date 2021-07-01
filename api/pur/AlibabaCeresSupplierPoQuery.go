package pur

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

/* AlibabaCeresSupplierPoQuery
采购供应商订单查询接口
alibaba.ceres.supplier.po.query

采购供应商订单查询接口 */
func AlibabaCeresSupplierPoQuery(clt *core.SDKClient, req *pur.AlibabaCeresSupplierPoQueryAPIRequest, session string) (*pur.AlibabaCeresSupplierPoQueryAPIResponse, error) {
	var resp pur.AlibabaCeresSupplierPoQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
