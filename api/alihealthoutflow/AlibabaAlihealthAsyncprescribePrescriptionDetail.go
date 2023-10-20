package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// Alibabaalihealthasyncprescribeprescriptiondetail 异步开方处方详情
// alibaba.alihealth.asyncprescribe.prescription.detail
//
// 异步开方处方查询
func Alibabaalihealthasyncprescribeprescriptiondetail(clt *core.SDKClient, req *alihealthoutflow.AlibabaalihealthasyncprescribeprescriptiondetailAPIRequest, session string) (*alihealthoutflow.AlibabaalihealthasyncprescribeprescriptiondetailAPIResponse, error) {
	var resp alihealthoutflow.AlibabaalihealthasyncprescribeprescriptiondetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
