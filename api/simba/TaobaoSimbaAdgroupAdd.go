package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbaadgroupadd 创建一个推广组
// taobao.simba.adgroup.add
//
// 创建一个推广组
func Taobaosimbaadgroupadd(clt *core.SDKClient, req *simba.TaobaosimbaadgroupaddAPIRequest, session string) (*simba.TaobaosimbaadgroupaddAPIResponse, error) {
	var resp simba.TaobaosimbaadgroupaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
