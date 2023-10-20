package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaAdgroupMobilediscountUpdate 对推广组进行单独移动溢价
// taobao.simba.adgroup.mobilediscount.update
//
// 对推广组进行单独移动溢价
func TaobaoSimbaAdgroupMobilediscountUpdate(clt *core.SDKClient, req *simba.TaobaoSimbaAdgroupMobilediscountUpdateAPIRequest, session string) (*simba.TaobaoSimbaAdgroupMobilediscountUpdateAPIResponse, error) {
	var resp simba.TaobaoSimbaAdgroupMobilediscountUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
