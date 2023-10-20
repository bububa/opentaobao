package ascpffo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

// Aliexpressascpffoitemquery AliExpress发货单明细分页查询API
// aliexpress.ascp.ffo.item.query
//
// AE履约发货单明细分页查询
func Aliexpressascpffoitemquery(clt *core.SDKClient, req *ascpffo.AliexpressascpffoitemqueryAPIRequest, session string) (*ascpffo.AliexpressascpffoitemqueryAPIResponse, error) {
	var resp ascpffo.AliexpressascpffoitemqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
