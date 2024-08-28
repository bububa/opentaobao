package wirelessshare

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wirelessshare"
)

// TaobaoWirelessShareTpwdQuery 查询解析淘口令
// taobao.wireless.share.tpwd.query
//
// 查询解析淘口令
func TaobaoWirelessShareTpwdQuery(ctx context.Context, clt *core.SDKClient, req *wirelessshare.TaobaoWirelessShareTpwdQueryAPIRequest, resp *wirelessshare.TaobaoWirelessShareTpwdQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
