package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminUserSuggest 获取关联账户列表
// yunos.tvpubadmin.user.suggest
//
// 获取关联账户列表
func YunosTvpubadminUserSuggest(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminUserSuggestAPIRequest, resp *tvupadmin.YunosTvpubadminUserSuggestAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
