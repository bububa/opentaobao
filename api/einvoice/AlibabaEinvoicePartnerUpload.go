package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoicepartnerupload 服务商发票上传接口（非授权）
// alibaba.einvoice.partner.upload
//
// 服务商发票上传接口（非授权）
func Alibabaeinvoicepartnerupload(clt *core.SDKClient, req *einvoice.AlibabaeinvoicepartneruploadAPIRequest, session string) (*einvoice.AlibabaeinvoicepartneruploadAPIResponse, error) {
	var resp einvoice.AlibabaeinvoicepartneruploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
