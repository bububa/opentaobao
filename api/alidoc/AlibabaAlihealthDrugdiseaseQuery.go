package alidoc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alidoc"
)

// AlibabaAlihealthDrugdiseaseQuery 药品诊断查询
// alibaba.alihealth.drugdisease.query
//
// 药品诊断查询
func AlibabaAlihealthDrugdiseaseQuery(clt *core.SDKClient, req *alidoc.AlibabaAlihealthDrugdiseaseQueryAPIRequest, session string) (*alidoc.AlibabaAlihealthDrugdiseaseQueryAPIResponse, error) {
	var resp alidoc.AlibabaAlihealthDrugdiseaseQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
