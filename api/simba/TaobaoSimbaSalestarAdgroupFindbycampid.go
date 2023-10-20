package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbasalestaradgroupfindbycampid (销量明星)批量获取推广计划下的推广组信息
// taobao.simba.salestar.adgroup.findbycampid
//
// 批量得到推广计划下的推广组
func Taobaosimbasalestaradgroupfindbycampid(clt *core.SDKClient, req *simba.TaobaosimbasalestaradgroupfindbycampidAPIRequest, session string) (*simba.TaobaosimbasalestaradgroupfindbycampidAPIResponse, error) {
	var resp simba.TaobaosimbasalestaradgroupfindbycampidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
