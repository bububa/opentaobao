package guoguo

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/guoguo"
)

// CainiaoGuoguoBackupGraborderTakepackage 兜底派送订单的揽件接口
// cainiao.guoguo.backup.graborder.takepackage
//
// 快递公司回传订单号和四位取件码给菜鸟裹裹
func CainiaoGuoguoBackupGraborderTakepackage(ctx context.Context, clt *core.SDKClient, req *guoguo.CainiaoGuoguoBackupGraborderTakepackageAPIRequest, resp *guoguo.CainiaoGuoguoBackupGraborderTakepackageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
