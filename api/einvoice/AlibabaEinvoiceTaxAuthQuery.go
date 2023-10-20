package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoicetaxauthquery 发票中台授权信息获取
// alibaba.einvoice.tax.auth.query
//
// 发票中台授权信息获取
func Alibabaeinvoicetaxauthquery(clt *core.SDKClient, req *einvoice.AlibabaeinvoicetaxauthqueryAPIRequest, session string) (*einvoice.AlibabaeinvoicetaxauthqueryAPIResponse, error) {
	var resp einvoice.AlibabaeinvoicetaxauthqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
