package ascpffo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

// AliexpressAscpPoItemQuery AliExpress采购单明细查询API
// aliexpress.ascp.po.item.query
//
// AE 供应链仓发 采购单明细查询
func AliexpressAscpPoItemQuery(clt *core.SDKClient, req *ascpffo.AliexpressAscpPoItemQueryAPIRequest, resp *ascpffo.AliexpressAscpPoItemQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
