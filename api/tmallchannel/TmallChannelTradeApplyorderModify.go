package tmallchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallchannel"
)

/* TmallChannelTradeApplyorderModify
供应商修改申请单
tmall.channel.trade.applyorder.modify

上游供应商修改申请单, 目前只允许修改价格+件数且sku数量必须完全一致 */
func TmallChannelTradeApplyorderModify(clt *core.SDKClient, req *tmallchannel.TmallChannelTradeApplyorderModifyAPIRequest, session string) (*tmallchannel.TmallChannelTradeApplyorderModifyAPIResponse, error) {
	var resp tmallchannel.TmallChannelTradeApplyorderModifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
