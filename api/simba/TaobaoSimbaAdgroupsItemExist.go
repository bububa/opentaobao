package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbaadgroupsitemexist 商品是否推广
// taobao.simba.adgroups.item.exist
//
// 判断在一个推广计划中是否已经推广了一个商品
func Taobaosimbaadgroupsitemexist(clt *core.SDKClient, req *simba.TaobaosimbaadgroupsitemexistAPIRequest, session string) (*simba.TaobaosimbaadgroupsitemexistAPIResponse, error) {
	var resp simba.TaobaosimbaadgroupsitemexistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
