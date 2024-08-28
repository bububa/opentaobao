package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSubwayWordpackageGet 获取词包列表
// taobao.subway.wordpackage.get
//
// 获取流量智选、捡漏词包等词包列表
func TaobaoSubwayWordpackageGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSubwayWordpackageGetAPIRequest, resp *simba.TaobaoSubwayWordpackageGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
