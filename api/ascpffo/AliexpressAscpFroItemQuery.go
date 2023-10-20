package ascpffo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

// Aliexpressascpfroitemquery AliExpress销退单明细查询API
// aliexpress.ascp.fro.item.query
//
// AE履约销退单明细查询API
func Aliexpressascpfroitemquery(clt *core.SDKClient, req *ascpffo.AliexpressascpfroitemqueryAPIRequest, session string) (*ascpffo.AliexpressascpfroitemqueryAPIResponse, error) {
	var resp ascpffo.AliexpressascpfroitemqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
