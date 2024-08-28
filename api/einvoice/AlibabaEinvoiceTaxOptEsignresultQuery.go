package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceTaxOptEsignresultQuery 查询用户签约税优结果
// alibaba.einvoice.tax.opt.esignresult.query
//
// 查询用户是否已经签约
func AlibabaEinvoiceTaxOptEsignresultQuery(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoiceTaxOptEsignresultQueryAPIRequest, resp *einvoice.AlibabaEinvoiceTaxOptEsignresultQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
