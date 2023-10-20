package guoguo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/guoguo"
)

// Cainiaoguoguobackupgrabordersubmitmailno 兜底派送订单的运单号回传接口
// cainiao.guoguo.backup.graborder.submitmailno
//
// 快递公司回传订单号和运单号给菜鸟裹裹
func Cainiaoguoguobackupgrabordersubmitmailno(clt *core.SDKClient, req *guoguo.CainiaoguoguobackupgrabordersubmitmailnoAPIRequest, session string) (*guoguo.CainiaoguoguobackupgrabordersubmitmailnoAPIResponse, error) {
	var resp guoguo.CainiaoguoguobackupgrabordersubmitmailnoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
