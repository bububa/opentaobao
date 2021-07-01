package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

/* TaobaoSimbaAdgroupidsChangedGet
获取修改的推广组ID
taobao.simba.adgroupids.changed.get

获取修改的推广组ID */
func TaobaoSimbaAdgroupidsChangedGet(clt *core.SDKClient, req *simba.TaobaoSimbaAdgroupidsChangedGetAPIRequest, session string) (*simba.TaobaoSimbaAdgroupidsChangedGetAPIResponse, error) {
	var resp simba.TaobaoSimbaAdgroupidsChangedGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
