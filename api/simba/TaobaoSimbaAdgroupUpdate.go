package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaAdgroupUpdate 更新一个推广组的信息
// taobao.simba.adgroup.update
//
// 更新一个推广组的信息，可以设置默认出价、是否上线、非搜索出价、非搜索是否使用默认出价
func TaobaoSimbaAdgroupUpdate(clt *core.SDKClient, req *simba.TaobaoSimbaAdgroupUpdateAPIRequest, session string) (*simba.TaobaoSimbaAdgroupUpdateAPIResponse, error) {
	var resp simba.TaobaoSimbaAdgroupUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
