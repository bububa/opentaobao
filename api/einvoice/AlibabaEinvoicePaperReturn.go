package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoicepaperreturn 纸质发票结果回传
// alibaba.einvoice.paper.return
//
// 纸质发票结果回传
func Alibabaeinvoicepaperreturn(clt *core.SDKClient, req *einvoice.AlibabaeinvoicepaperreturnAPIRequest, session string) (*einvoice.AlibabaeinvoicepaperreturnAPIResponse, error) {
	var resp einvoice.AlibabaeinvoicepaperreturnAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
