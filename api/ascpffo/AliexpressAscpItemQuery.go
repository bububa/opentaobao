package ascpffo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

// AliexpressAscpItemQuery AliExpress货品查询查询API
// aliexpress.ascp.item.query
//
// AE货品查询API
func AliexpressAscpItemQuery(clt *core.SDKClient, req *ascpffo.AliexpressAscpItemQueryAPIRequest, resp *ascpffo.AliexpressAscpItemQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
