package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSubwayAccountOfflineFind 获取账户历史多日汇总报表
// taobao.subway.account.offline.find
//
// 获取账户历史报表
func TaobaoSubwayAccountOfflineFind(clt *core.SDKClient, req *simba.TaobaoSubwayAccountOfflineFindAPIRequest, session string) (*simba.TaobaoSubwayAccountOfflineFindAPIResponse, error) {
	var resp simba.TaobaoSubwayAccountOfflineFindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
