package guoguo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/guoguo"
)

// CainiaoGuoguoBackupGraborderTakepackage 兜底派送订单的揽件接口
// cainiao.guoguo.backup.graborder.takepackage
//
// 快递公司回传订单号和四位取件码给菜鸟裹裹
func CainiaoGuoguoBackupGraborderTakepackage(clt *core.SDKClient, req *guoguo.CainiaoGuoguoBackupGraborderTakepackageAPIRequest, session string) (*guoguo.CainiaoGuoguoBackupGraborderTakepackageAPIResponse, error) {
	var resp guoguo.CainiaoGuoguoBackupGraborderTakepackageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
