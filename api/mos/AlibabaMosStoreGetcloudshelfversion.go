package mos

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMosStoreGetcloudshelfversion 获取云货架版本信息
// alibaba.mos.store.getcloudshelfversion
//
// 根据屏编号获取云货架版本信息
func AlibabaMosStoreGetcloudshelfversion(ctx context.Context, clt *core.SDKClient, req *mos.AlibabaMosStoreGetcloudshelfversionAPIRequest, resp *mos.AlibabaMosStoreGetcloudshelfversionAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
