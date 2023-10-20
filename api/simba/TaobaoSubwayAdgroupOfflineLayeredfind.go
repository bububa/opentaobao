package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSubwayAdgroupOfflineLayeredfind 查询单元离线列表30天转化周期
// taobao.subway.adgroup.offline.layeredfind
//
// 查询单元离线列表
func TaobaoSubwayAdgroupOfflineLayeredfind(clt *core.SDKClient, req *simba.TaobaoSubwayAdgroupOfflineLayeredfindAPIRequest, session string) (*simba.TaobaoSubwayAdgroupOfflineLayeredfindAPIResponse, error) {
	var resp simba.TaobaoSubwayAdgroupOfflineLayeredfindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
