package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSubwayAccountOfflineLayeredfind 获取账户历史报表30天转化周期
// taobao.subway.account.offline.layeredfind
//
// 获取账户历史报表
func TaobaoSubwayAccountOfflineLayeredfind(clt *core.SDKClient, req *simba.TaobaoSubwayAccountOfflineLayeredfindAPIRequest, session string) (*simba.TaobaoSubwayAccountOfflineLayeredfindAPIResponse, error) {
	var resp simba.TaobaoSubwayAccountOfflineLayeredfindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
