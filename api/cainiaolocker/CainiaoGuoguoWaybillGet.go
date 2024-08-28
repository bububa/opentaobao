package cainiaolocker

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaolocker"
)

// CainiaoGuoguoWaybillGet 菜鸟裹裹商家寄件取号接口
// cainiao.guoguo.waybill.get
//
// 菜鸟裹裹商家寄件取号接口
func CainiaoGuoguoWaybillGet(ctx context.Context, clt *core.SDKClient, req *cainiaolocker.CainiaoGuoguoWaybillGetAPIRequest, resp *cainiaolocker.CainiaoGuoguoWaybillGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
