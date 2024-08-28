package cainiaoecc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaoecc"
)

// CainiaoEccExceptionsDelayGet 菜鸟控制塔包裹滞留异常信息获取
// cainiao.ecc.exceptions.delay.get
//
// 菜鸟控制塔包裹滞留异常信息获取
func CainiaoEccExceptionsDelayGet(ctx context.Context, clt *core.SDKClient, req *cainiaoecc.CainiaoEccExceptionsDelayGetAPIRequest, resp *cainiaoecc.CainiaoEccExceptionsDelayGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
