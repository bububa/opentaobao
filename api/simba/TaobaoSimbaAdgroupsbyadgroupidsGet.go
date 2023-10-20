package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaAdgroupsbyadgroupidsGet 批量得到推广组
// taobao.simba.adgroupsbyadgroupids.get
//
// 批量得到推广组
func TaobaoSimbaAdgroupsbyadgroupidsGet(clt *core.SDKClient, req *simba.TaobaoSimbaAdgroupsbyadgroupidsGetAPIRequest, session string) (*simba.TaobaoSimbaAdgroupsbyadgroupidsGetAPIResponse, error) {
	var resp simba.TaobaoSimbaAdgroupsbyadgroupidsGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
