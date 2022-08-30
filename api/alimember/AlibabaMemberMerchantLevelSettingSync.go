package alimember

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimember"
)

// AlibabaMemberMerchantLevelSettingSync 商家等级列表同步配置
// alibaba.member.merchant.level.setting.sync
//
// 商家等级列表同步配置
func AlibabaMemberMerchantLevelSettingSync(clt *core.SDKClient, req *alimember.AlibabaMemberMerchantLevelSettingSyncAPIRequest, session string) (*alimember.AlibabaMemberMerchantLevelSettingSyncAPIResponse, error) {
	var resp alimember.AlibabaMemberMerchantLevelSettingSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
