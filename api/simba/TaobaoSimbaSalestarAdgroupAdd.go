package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbasalestaradgroupadd (新)创建一个推广组
// taobao.simba.salestar.adgroup.add
//
// 创建一个推广组
func Taobaosimbasalestaradgroupadd(clt *core.SDKClient, req *simba.TaobaosimbasalestaradgroupaddAPIRequest, session string) (*simba.TaobaosimbasalestaradgroupaddAPIResponse, error) {
	var resp simba.TaobaosimbasalestaradgroupaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
