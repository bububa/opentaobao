package ascpffo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

// Aliexpressascpinventorylogquery AliExpress库存流水查询API
// aliexpress.ascp.inventory.log.query
//
// AliExpress库存流水查询API
func Aliexpressascpinventorylogquery(clt *core.SDKClient, req *ascpffo.AliexpressascpinventorylogqueryAPIRequest, session string) (*ascpffo.AliexpressascpinventorylogqueryAPIResponse, error) {
	var resp ascpffo.AliexpressascpinventorylogqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
