package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbarptcustbaseget 客户账户报表基础数据对象
// taobao.simba.rpt.custbase.get
//
// 客户账户报表基础数据对象
func Taobaosimbarptcustbaseget(clt *core.SDKClient, req *simba.TaobaosimbarptcustbasegetAPIRequest, session string) (*simba.TaobaosimbarptcustbasegetAPIResponse, error) {
	var resp simba.TaobaosimbarptcustbasegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
