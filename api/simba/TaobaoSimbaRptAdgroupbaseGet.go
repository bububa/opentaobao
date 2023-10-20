package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbarptadgroupbaseget 推广组基础报表数据对象
// taobao.simba.rpt.adgroupbase.get
//
// 推广组基础报表数据对象
func Taobaosimbarptadgroupbaseget(clt *core.SDKClient, req *simba.TaobaosimbarptadgroupbasegetAPIRequest, session string) (*simba.TaobaosimbarptadgroupbasegetAPIResponse, error) {
	var resp simba.TaobaosimbarptadgroupbasegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
