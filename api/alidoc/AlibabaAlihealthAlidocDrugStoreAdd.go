package alidoc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alidoc"
)

// Alibabaalihealthalidocdrugstoreadd gsk新增药店
// alibaba.alihealth.alidoc.drug.store.add
//
// GSK上传药店信息
func Alibabaalihealthalidocdrugstoreadd(clt *core.SDKClient, req *alidoc.AlibabaalihealthalidocdrugstoreaddAPIRequest, session string) (*alidoc.AlibabaalihealthalidocdrugstoreaddAPIResponse, error) {
	var resp alidoc.AlibabaalihealthalidocdrugstoreaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
