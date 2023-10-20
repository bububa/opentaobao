package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbartrptcustget 获取账户实时报表数据
// taobao.simba.rtrpt.cust.get
//
// 获取账户实时报表数据
func Taobaosimbartrptcustget(clt *core.SDKClient, req *simba.TaobaosimbartrptcustgetAPIRequest, session string) (*simba.TaobaosimbartrptcustgetAPIResponse, error) {
	var resp simba.TaobaosimbartrptcustgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
