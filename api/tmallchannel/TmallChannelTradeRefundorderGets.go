package tmallchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallchannel"
)

// Tmallchanneltraderefundordergets 供应商查询退款单
// tmall.channel.trade.refundorder.gets
//
// 供应商分页查询退款单
func Tmallchanneltraderefundordergets(clt *core.SDKClient, req *tmallchannel.TmallchanneltraderefundordergetsAPIRequest, session string) (*tmallchannel.TmallchanneltraderefundordergetsAPIResponse, error) {
	var resp tmallchannel.TmallchanneltraderefundordergetsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
