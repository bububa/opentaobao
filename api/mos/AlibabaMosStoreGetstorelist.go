package mos

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMosStoreGetstorelist 根据屏编号获取专柜集
// alibaba.mos.store.getstorelist
//
// 根据屏编号获取专柜集
func AlibabaMosStoreGetstorelist(ctx context.Context, clt *core.SDKClient, req *mos.AlibabaMosStoreGetstorelistAPIRequest, resp *mos.AlibabaMosStoreGetstorelistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
