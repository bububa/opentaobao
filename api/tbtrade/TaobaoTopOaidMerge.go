package tbtrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// TaobaoTopOaidMerge OAID订单合并
// taobao.top.oaid.merge
//
// 基于OAID（收件人ID， Open Addressee ID)做订单合并，确保相同收件人信息的订单合并到相同组。
func TaobaoTopOaidMerge(clt *core.SDKClient, req *tbtrade.TaobaoTopOaidMergeAPIRequest, session string) (*tbtrade.TaobaoTopOaidMergeAPIResponse, error) {
	var resp tbtrade.TaobaoTopOaidMergeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
