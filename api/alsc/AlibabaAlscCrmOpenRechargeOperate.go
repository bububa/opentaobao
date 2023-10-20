package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmopenrechargeoperate 储值操作接口
// alibaba.alsc.crm.open.recharge.operate
//
// 储值操作接口
func Alibabaalsccrmopenrechargeoperate(clt *core.SDKClient, req *alsc.AlibabaalsccrmopenrechargeoperateAPIRequest, session string) (*alsc.AlibabaalsccrmopenrechargeoperateAPIResponse, error) {
	var resp alsc.AlibabaalsccrmopenrechargeoperateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
