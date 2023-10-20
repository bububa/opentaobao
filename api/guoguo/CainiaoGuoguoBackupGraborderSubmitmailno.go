package guoguo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/guoguo"
)

// CainiaoGuoguoBackupGraborderSubmitmailno 兜底派送订单的运单号回传接口
// cainiao.guoguo.backup.graborder.submitmailno
//
// 快递公司回传订单号和运单号给菜鸟裹裹
func CainiaoGuoguoBackupGraborderSubmitmailno(clt *core.SDKClient, req *guoguo.CainiaoGuoguoBackupGraborderSubmitmailnoAPIRequest, resp *guoguo.CainiaoGuoguoBackupGraborderSubmitmailnoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
