package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminUserRights 获取用户权益
// yunos.tvpubadmin.user.rights
//
// 获取用户权益
func YunosTvpubadminUserRights(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminUserRightsAPIRequest, resp *tvupadmin.YunosTvpubadminUserRightsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
