package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthAsyncprescribePrescriptionSearch 异步开方处方查询
// alibaba.alihealth.asyncprescribe.prescription.search
//
// 异步开方处方查询
func AlibabaAlihealthAsyncprescribePrescriptionSearch(clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthAsyncprescribePrescriptionSearchAPIRequest, resp *alihealthoutflow.AlibabaAlihealthAsyncprescribePrescriptionSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
