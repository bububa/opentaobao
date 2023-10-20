package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Taobaoinvturnoverquery 业务库存流水查询
// taobao.inv.turnover.query
//
// 业务库存流水
func Taobaoinvturnoverquery(clt *core.SDKClient, req *ascpchannel.TaobaoinvturnoverqueryAPIRequest, session string) (*ascpchannel.TaobaoinvturnoverqueryAPIResponse, error) {
	var resp ascpchannel.TaobaoinvturnoverqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
