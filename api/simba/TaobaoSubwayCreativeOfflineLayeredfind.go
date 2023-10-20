package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSubwayCreativeOfflineLayeredfind 获取创意离线报表周期30天
// taobao.subway.creative.offline.layeredfind
//
// 获取创意离线报表
func TaobaoSubwayCreativeOfflineLayeredfind(clt *core.SDKClient, req *simba.TaobaoSubwayCreativeOfflineLayeredfindAPIRequest, session string) (*simba.TaobaoSubwayCreativeOfflineLayeredfindAPIResponse, error) {
	var resp simba.TaobaoSubwayCreativeOfflineLayeredfindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
