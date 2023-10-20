package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// Alibabaalihealthasyncprescribeprescriptionsearch 异步开方处方查询
// alibaba.alihealth.asyncprescribe.prescription.search
//
// 异步开方处方查询
func Alibabaalihealthasyncprescribeprescriptionsearch(clt *core.SDKClient, req *alihealthoutflow.AlibabaalihealthasyncprescribeprescriptionsearchAPIRequest, session string) (*alihealthoutflow.AlibabaalihealthasyncprescribeprescriptionsearchAPIResponse, error) {
	var resp alihealthoutflow.AlibabaalihealthasyncprescribeprescriptionsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
