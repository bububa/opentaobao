package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoiceflowrenew 工单(入驻、加盘、续约)续约
// alibaba.einvoice.flow.renew
//
// 工单(含入驻、加盘、续约工单)续约能力开放
func Alibabaeinvoiceflowrenew(clt *core.SDKClient, req *einvoice.AlibabaeinvoiceflowrenewAPIRequest, session string) (*einvoice.AlibabaeinvoiceflowrenewAPIResponse, error) {
	var resp einvoice.AlibabaeinvoiceflowrenewAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
