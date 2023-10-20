package alimember

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimember"
)

// Alibabamembermerchantlevelsettingsync 商家等级列表同步配置
// alibaba.member.merchant.level.setting.sync
//
// 商家等级列表同步配置
func Alibabamembermerchantlevelsettingsync(clt *core.SDKClient, req *alimember.AlibabamembermerchantlevelsettingsyncAPIRequest, session string) (*alimember.AlibabamembermerchantlevelsettingsyncAPIResponse, error) {
	var resp alimember.AlibabamembermerchantlevelsettingsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
