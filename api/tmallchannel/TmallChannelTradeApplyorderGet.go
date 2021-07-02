package tmallchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallchannel"
)

// TmallChannelTradeApplyorderGet 查询采购申请单详情
// tmall.channel.trade.applyorder.get
//
// 通过采购申请单ID获取单据详情
func TmallChannelTradeApplyorderGet(clt *core.SDKClient, req *tmallchannel.TmallChannelTradeApplyorderGetAPIRequest, session string) (*tmallchannel.TmallChannelTradeApplyorderGetAPIResponse, error) {
	var resp tmallchannel.TmallChannelTradeApplyorderGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
