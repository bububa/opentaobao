package ascpffo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

// AliexpressAscpPoQuery AliExpress采购单查询API
// aliexpress.ascp.po.query
//
// AE仓发业务采购单查询
func AliexpressAscpPoQuery(clt *core.SDKClient, req *ascpffo.AliexpressAscpPoQueryAPIRequest, resp *ascpffo.AliexpressAscpPoQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
