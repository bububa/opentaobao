package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoicetaxoptesignresultquery 查询用户签约税优结果
// alibaba.einvoice.tax.opt.esignresult.query
//
// 查询用户是否已经签约
func Alibabaeinvoicetaxoptesignresultquery(clt *core.SDKClient, req *einvoice.AlibabaeinvoicetaxoptesignresultqueryAPIRequest, session string) (*einvoice.AlibabaeinvoicetaxoptesignresultqueryAPIResponse, error) {
	var resp einvoice.AlibabaeinvoicetaxoptesignresultqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
