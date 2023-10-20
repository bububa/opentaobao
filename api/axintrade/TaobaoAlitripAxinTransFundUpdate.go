package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// Taobaoalitripaxintransfundupdate 修改资金单接口
// taobao.alitrip.axin.trans.fund.update
//
// 阿信供销平台-修改资金单接口
func Taobaoalitripaxintransfundupdate(clt *core.SDKClient, req *axintrade.TaobaoalitripaxintransfundupdateAPIRequest, session string) (*axintrade.TaobaoalitripaxintransfundupdateAPIResponse, error) {
	var resp axintrade.TaobaoalitripaxintransfundupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
