package util

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoTopIpoutGet 获取开放平台出口IP段
// taobao.top.ipout.get
//
// 获取开放平台出口IP段
func TaobaoTopIpoutGet(ctx context.Context, clt *core.SDKClient, req *util.TaobaoTopIpoutGetAPIRequest, resp *util.TaobaoTopIpoutGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
