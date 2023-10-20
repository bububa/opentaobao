package tmallchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallchannel"
)

// Tmallchanneltradeorderget 通过主采购单号查询采购单
// tmall.channel.trade.order.get
//
// 通过主采购单号查询采购单
func Tmallchanneltradeorderget(clt *core.SDKClient, req *tmallchannel.TmallchanneltradeordergetAPIRequest, session string) (*tmallchannel.TmallchanneltradeordergetAPIResponse, error) {
	var resp tmallchannel.TmallchanneltradeordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
