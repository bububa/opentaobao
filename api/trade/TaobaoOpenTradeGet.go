package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// Taobaoopentradeget 获取单笔交易的部分信息(商家应用使用)
// taobao.open.trade.get
//
// 获取单笔交易的部分信息&lt;/br&gt;
// 1.入参fields中传入buyer_nick ，才能返回buyer_open_id
func Taobaoopentradeget(clt *core.SDKClient, req *trade.TaobaoopentradegetAPIRequest, session string) (*trade.TaobaoopentradegetAPIResponse, error) {
	var resp trade.TaobaoopentradegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
