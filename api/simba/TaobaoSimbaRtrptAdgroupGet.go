package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbartrptadgroupget 获取推广组实时报表数据
// taobao.simba.rtrpt.adgroup.get
//
// 获取推广组实时报表数据
func Taobaosimbartrptadgroupget(clt *core.SDKClient, req *simba.TaobaosimbartrptadgroupgetAPIRequest, session string) (*simba.TaobaosimbartrptadgroupgetAPIResponse, error) {
	var resp simba.TaobaosimbartrptadgroupgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
