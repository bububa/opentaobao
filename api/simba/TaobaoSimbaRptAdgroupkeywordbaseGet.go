package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbarptadgroupkeywordbaseget 推广组下的词基础报表数据查询(明细数据不分类型查询)
// taobao.simba.rpt.adgroupkeywordbase.get
//
// 推广组下的词基础报表数据查询(明细数据不分类型查询)
func Taobaosimbarptadgroupkeywordbaseget(clt *core.SDKClient, req *simba.TaobaosimbarptadgroupkeywordbasegetAPIRequest, session string) (*simba.TaobaosimbarptadgroupkeywordbasegetAPIResponse, error) {
	var resp simba.TaobaosimbarptadgroupkeywordbasegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
