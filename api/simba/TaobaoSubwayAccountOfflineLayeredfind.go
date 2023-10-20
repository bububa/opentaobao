package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosubwayaccountofflinelayeredfind 获取账户历史报表30天转化周期
// taobao.subway.account.offline.layeredfind
//
// 获取账户历史报表
func Taobaosubwayaccountofflinelayeredfind(clt *core.SDKClient, req *simba.TaobaosubwayaccountofflinelayeredfindAPIRequest, session string) (*simba.TaobaosubwayaccountofflinelayeredfindAPIResponse, error) {
	var resp simba.TaobaosubwayaccountofflinelayeredfindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
