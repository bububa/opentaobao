package jym

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// AlibabaJymIndustryOutsidegamepropertysyncSyncpropertyinfo 外部上报游戏属性信息
// alibaba.jym.industry.outsidegamepropertysync.syncpropertyinfo
//
// 外部上报游戏属性信息
func AlibabaJymIndustryOutsidegamepropertysyncSyncpropertyinfo(ctx context.Context, clt *core.SDKClient, req *jym.AlibabaJymIndustryOutsidegamepropertysyncSyncpropertyinfoAPIRequest, resp *jym.AlibabaJymIndustryOutsidegamepropertysyncSyncpropertyinfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
