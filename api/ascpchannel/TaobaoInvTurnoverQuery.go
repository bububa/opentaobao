package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// TaobaoInvTurnoverQuery 业务库存流水查询
// taobao.inv.turnover.query
//
// 业务库存流水
func TaobaoInvTurnoverQuery(clt *core.SDKClient, req *ascpchannel.TaobaoInvTurnoverQueryAPIRequest, resp *ascpchannel.TaobaoInvTurnoverQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
