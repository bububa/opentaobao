package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoicepaperprint 纸票打印接口
// alibaba.einvoice.paper.print
//
// 打印一张已开具成功的纸票
func Alibabaeinvoicepaperprint(clt *core.SDKClient, req *einvoice.AlibabaeinvoicepaperprintAPIRequest, session string) (*einvoice.AlibabaeinvoicepaperprintAPIResponse, error) {
	var resp einvoice.AlibabaeinvoicepaperprintAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
