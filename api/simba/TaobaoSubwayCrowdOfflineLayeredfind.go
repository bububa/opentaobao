package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSubwayCrowdOfflineLayeredfind 获取人群离线报表30转化周期
// taobao.subway.crowd.offline.layeredfind
//
// 获取人群离线报表
func TaobaoSubwayCrowdOfflineLayeredfind(clt *core.SDKClient, req *simba.TaobaoSubwayCrowdOfflineLayeredfindAPIRequest, session string) (*simba.TaobaoSubwayCrowdOfflineLayeredfindAPIResponse, error) {
	var resp simba.TaobaoSubwayCrowdOfflineLayeredfindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
