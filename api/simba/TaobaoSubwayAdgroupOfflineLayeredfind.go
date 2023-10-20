package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSubwayAdgroupOfflineLayeredfind 查询单元离线列表30天转化周期
// taobao.subway.adgroup.offline.layeredfind
//
// 查询单元离线列表
func TaobaoSubwayAdgroupOfflineLayeredfind(clt *core.SDKClient, req *simba.TaobaoSubwayAdgroupOfflineLayeredfindAPIRequest, resp *simba.TaobaoSubwayAdgroupOfflineLayeredfindAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
