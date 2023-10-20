package pur

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// Alibabapursupplierasncreate asn创建
// alibaba.pur.supplier.asncreate
//
// asn创建
func Alibabapursupplierasncreate(clt *core.SDKClient, req *pur.AlibabapursupplierasncreateAPIRequest, session string) (*pur.AlibabapursupplierasncreateAPIResponse, error) {
	var resp pur.AlibabapursupplierasncreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
