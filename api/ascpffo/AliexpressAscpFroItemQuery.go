package ascpffo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

// AliexpressAscpFroItemQuery AliExpress销退单明细查询API
// aliexpress.ascp.fro.item.query
//
// AE履约销退单明细查询API
func AliexpressAscpFroItemQuery(clt *core.SDKClient, req *ascpffo.AliexpressAscpFroItemQueryAPIRequest, resp *ascpffo.AliexpressAscpFroItemQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
