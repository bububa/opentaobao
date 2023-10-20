package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosubwayaccountofflinefind 获取账户历史多日汇总报表
// taobao.subway.account.offline.find
//
// 获取账户历史报表
func Taobaosubwayaccountofflinefind(clt *core.SDKClient, req *simba.TaobaosubwayaccountofflinefindAPIRequest, session string) (*simba.TaobaosubwayaccountofflinefindAPIResponse, error) {
	var resp simba.TaobaosubwayaccountofflinefindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
