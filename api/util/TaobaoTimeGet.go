package util

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoTimeGet 获取淘宝系统当前时间
// taobao.time.get
//
// 获取淘宝系统当前时间
func TaobaoTimeGet(ctx context.Context, clt *core.SDKClient, req *util.TaobaoTimeGetAPIRequest, resp *util.TaobaoTimeGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
