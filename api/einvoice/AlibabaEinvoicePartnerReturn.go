package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoicepartnerreturn 开票商回传开票结果
// alibaba.einvoice.partner.return
//
// 开票商返回开票结果数据
func Alibabaeinvoicepartnerreturn(clt *core.SDKClient, req *einvoice.AlibabaeinvoicepartnerreturnAPIRequest, session string) (*einvoice.AlibabaeinvoicepartnerreturnAPIResponse, error) {
	var resp einvoice.AlibabaeinvoicepartnerreturnAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
