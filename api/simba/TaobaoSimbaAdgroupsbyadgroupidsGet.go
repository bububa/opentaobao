package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbaadgroupsbyadgroupidsget 批量得到推广组
// taobao.simba.adgroupsbyadgroupids.get
//
// 批量得到推广组
func Taobaosimbaadgroupsbyadgroupidsget(clt *core.SDKClient, req *simba.TaobaosimbaadgroupsbyadgroupidsgetAPIRequest, session string) (*simba.TaobaosimbaadgroupsbyadgroupidsgetAPIResponse, error) {
	var resp simba.TaobaosimbaadgroupsbyadgroupidsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
