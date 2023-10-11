package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsExpressUserBlacklistTmsSync 上门取退用户黑名单同步
// taobao.logistics.express.user.blacklist.tms.sync
//
// 上门取退用户黑名单同步
func TaobaoLogisticsExpressUserBlacklistTmsSync(clt *core.SDKClient, req *ascp.TaobaoLogisticsExpressUserBlacklistTmsSyncAPIRequest, session string) (*ascp.TaobaoLogisticsExpressUserBlacklistTmsSyncAPIResponse, error) {
	var resp ascp.TaobaoLogisticsExpressUserBlacklistTmsSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
