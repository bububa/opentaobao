package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpAccountIsUniversalUser 判断用户是否迁移新bp
// taobao.universalbp.account.is.universal.user
//
// 获取客户是否应使用新接口。对于迁移了新bp的客户，使用新接口，没有迁移的，使用老bp接口。不可错乱使用。
func TaobaoUniversalbpAccountIsUniversalUser(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpAccountIsUniversalUserAPIRequest, resp *simba.TaobaoUniversalbpAccountIsUniversalUserAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
