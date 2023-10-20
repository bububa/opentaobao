package alidoc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alidoc"
)

// Alibabaalihealthdrugdiseasequery 药品诊断查询
// alibaba.alihealth.drugdisease.query
//
// 药品诊断查询
func Alibabaalihealthdrugdiseasequery(clt *core.SDKClient, req *alidoc.AlibabaalihealthdrugdiseasequeryAPIRequest, session string) (*alidoc.AlibabaalihealthdrugdiseasequeryAPIResponse, error) {
	var resp alidoc.AlibabaalihealthdrugdiseasequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
