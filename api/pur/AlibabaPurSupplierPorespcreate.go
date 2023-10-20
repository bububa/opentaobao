package pur

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// Alibabapursupplierporespcreate po反馈创建
// alibaba.pur.supplier.porespcreate
//
// PO反馈接口
func Alibabapursupplierporespcreate(clt *core.SDKClient, req *pur.AlibabapursupplierporespcreateAPIRequest, session string) (*pur.AlibabapursupplierporespcreateAPIResponse, error) {
	var resp pur.AlibabapursupplierporespcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
