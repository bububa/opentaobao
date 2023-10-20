package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// Alibabaalihealthrxcaprescribesignedstatussave 处方ca认证
// alibaba.alihealth.rx.ca.prescribe.signed.status.save
//
// 处方ca认证
func Alibabaalihealthrxcaprescribesignedstatussave(clt *core.SDKClient, req *alihealthoutflow.AlibabaalihealthrxcaprescribesignedstatussaveAPIRequest, session string) (*alihealthoutflow.AlibabaalihealthrxcaprescribesignedstatussaveAPIResponse, error) {
	var resp alihealthoutflow.AlibabaalihealthrxcaprescribesignedstatussaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
