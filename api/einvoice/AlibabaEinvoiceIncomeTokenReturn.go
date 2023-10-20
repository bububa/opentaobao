package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoiceincometokenreturn 服务商回传税号token
// alibaba.einvoice.income.token.return
//
// 服务商回传税号token，用来勾选抵扣认证
func Alibabaeinvoiceincometokenreturn(clt *core.SDKClient, req *einvoice.AlibabaeinvoiceincometokenreturnAPIRequest, session string) (*einvoice.AlibabaeinvoiceincometokenreturnAPIResponse, error) {
	var resp einvoice.AlibabaeinvoiceincometokenreturnAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
