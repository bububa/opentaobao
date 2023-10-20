package traveltrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// Taobaoalitriptraveltradememoupdate 修改一笔交易备注
// taobao.alitrip.travel.trade.memo.update
//
// 更新一笔交易备注
func Taobaoalitriptraveltradememoupdate(clt *core.SDKClient, req *traveltrade.TaobaoalitriptraveltradememoupdateAPIRequest, session string) (*traveltrade.TaobaoalitriptraveltradememoupdateAPIResponse, error) {
	var resp traveltrade.TaobaoalitriptraveltradememoupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
