package pur

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// Alibabapursupplierinvoicecreate preInvoice创建
// alibaba.pur.supplier.invoicecreate
//
// preInvoice创建
func Alibabapursupplierinvoicecreate(clt *core.SDKClient, req *pur.AlibabapursupplierinvoicecreateAPIRequest, session string) (*pur.AlibabapursupplierinvoicecreateAPIResponse, error) {
	var resp pur.AlibabapursupplierinvoicecreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
