package traveltrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// Taobaoalitriptraveltradememoadd 添加一笔交易备注
// taobao.alitrip.travel.trade.memo.add
//
// 对一笔交易添加备注
func Taobaoalitriptraveltradememoadd(clt *core.SDKClient, req *traveltrade.TaobaoalitriptraveltradememoaddAPIRequest, session string) (*traveltrade.TaobaoalitriptraveltradememoaddAPIResponse, error) {
	var resp traveltrade.TaobaoalitriptraveltradememoaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
