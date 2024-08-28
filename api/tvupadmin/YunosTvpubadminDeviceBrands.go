package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceBrands 获取终端类型下品牌列表
// yunos.tvpubadmin.device.brands
//
// 获取终端类型下品牌列表
func YunosTvpubadminDeviceBrands(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceBrandsAPIRequest, resp *tvupadmin.YunosTvpubadminDeviceBrandsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
