package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// Taobaoalitripaxintransrefundcreate 阿信创建退款单
// taobao.alitrip.axin.trans.refund.create
//
// 阿信供销平台-创建退款单服务
func Taobaoalitripaxintransrefundcreate(clt *core.SDKClient, req *axintrade.TaobaoalitripaxintransrefundcreateAPIRequest, session string) (*axintrade.TaobaoalitripaxintransrefundcreateAPIResponse, error) {
	var resp axintrade.TaobaoalitripaxintransrefundcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
