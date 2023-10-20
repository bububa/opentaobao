package guoguo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/guoguo"
)

// Cainiaoguoguobackupgrabordertakepackage 兜底派送订单的揽件接口
// cainiao.guoguo.backup.graborder.takepackage
//
// 快递公司回传订单号和四位取件码给菜鸟裹裹
func Cainiaoguoguobackupgrabordertakepackage(clt *core.SDKClient, req *guoguo.CainiaoguoguobackupgrabordertakepackageAPIRequest, session string) (*guoguo.CainiaoguoguobackupgrabordertakepackageAPIResponse, error) {
	var resp guoguo.CainiaoguoguobackupgrabordertakepackageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
