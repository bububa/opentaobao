package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsExpressUserBlacklistTmsSync 上门取退用户黑名单同步
// taobao.logistics.express.user.blacklist.tms.sync
//
// 上门取退用户黑名单同步
func TaobaoLogisticsExpressUserBlacklistTmsSync(clt *core.SDKClient, req *ascp.TaobaoLogisticsExpressUserBlacklistTmsSyncAPIRequest, resp *ascp.TaobaoLogisticsExpressUserBlacklistTmsSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
