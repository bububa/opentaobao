package alimember

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimember"
)

// AlibabaMemberMerchantLevelSettingSync 商家等级列表同步配置
// alibaba.member.merchant.level.setting.sync
//
// 商家等级列表同步配置
func AlibabaMemberMerchantLevelSettingSync(ctx context.Context, clt *core.SDKClient, req *alimember.AlibabaMemberMerchantLevelSettingSyncAPIRequest, resp *alimember.AlibabaMemberMerchantLevelSettingSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
