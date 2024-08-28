package shenjing

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shenjing"
)

// AlibabaShenjingCoreActivityGetappshowlist 获取神鲸活动列表
// alibaba.shenjing.core.activity.getappshowlist
//
// 获取神鲸活动列表
func AlibabaShenjingCoreActivityGetappshowlist(ctx context.Context, clt *core.SDKClient, req *shenjing.AlibabaShenjingCoreActivityGetappshowlistAPIRequest, resp *shenjing.AlibabaShenjingCoreActivityGetappshowlistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
