package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbarptadgroupeffectget 推广组效果报表数据对象
// taobao.simba.rpt.adgroupeffect.get
//
// 推广组效果报表数据对象
func Taobaosimbarptadgroupeffectget(clt *core.SDKClient, req *simba.TaobaosimbarptadgroupeffectgetAPIRequest, session string) (*simba.TaobaosimbarptadgroupeffectgetAPIResponse, error) {
	var resp simba.TaobaosimbarptadgroupeffectgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
