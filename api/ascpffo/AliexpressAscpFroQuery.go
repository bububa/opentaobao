package ascpffo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

// AliexpressAscpFroQuery AliExpress销退单查询API
// aliexpress.ascp.fro.query
//
// AE履约销退单查询接口
func AliexpressAscpFroQuery(clt *core.SDKClient, req *ascpffo.AliexpressAscpFroQueryAPIRequest, resp *ascpffo.AliexpressAscpFroQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
