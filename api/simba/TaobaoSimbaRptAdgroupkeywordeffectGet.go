package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbarptadgroupkeywordeffectget 推广组下的词效果报表数据查询(明细数据不分类型查询)
// taobao.simba.rpt.adgroupkeywordeffect.get
//
// 推广组下的词效果报表数据查询(明细数据不分类型查询)
func Taobaosimbarptadgroupkeywordeffectget(clt *core.SDKClient, req *simba.TaobaosimbarptadgroupkeywordeffectgetAPIRequest, session string) (*simba.TaobaosimbarptadgroupkeywordeffectgetAPIResponse, error) {
	var resp simba.TaobaosimbarptadgroupkeywordeffectgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
