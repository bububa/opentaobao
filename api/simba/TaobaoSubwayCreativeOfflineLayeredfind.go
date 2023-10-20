package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSubwayCreativeOfflineLayeredfind 获取创意离线报表周期30天
// taobao.subway.creative.offline.layeredfind
//
// 获取创意离线报表
func TaobaoSubwayCreativeOfflineLayeredfind(clt *core.SDKClient, req *simba.TaobaoSubwayCreativeOfflineLayeredfindAPIRequest, resp *simba.TaobaoSubwayCreativeOfflineLayeredfindAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
