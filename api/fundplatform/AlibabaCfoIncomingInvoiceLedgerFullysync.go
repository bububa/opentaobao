package fundplatform

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

// Alibabacfoincominginvoiceledgerfullysync 票易通全量底账数据同步
// alibaba.cfo.incoming.invoice.ledger.fullysync
//
// 票易通全量底账数据同步
func Alibabacfoincominginvoiceledgerfullysync(clt *core.SDKClient, req *fundplatform.AlibabacfoincominginvoiceledgerfullysyncAPIRequest, session string) (*fundplatform.AlibabacfoincominginvoiceledgerfullysyncAPIResponse, error) {
	var resp fundplatform.AlibabacfoincominginvoiceledgerfullysyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
