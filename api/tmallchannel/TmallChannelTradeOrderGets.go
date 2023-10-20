package tmallchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallchannel"
)

// Tmallchanneltradeordergets 分页查询采购单
// tmall.channel.trade.order.gets
//
// 分页查询采购单
func Tmallchanneltradeordergets(clt *core.SDKClient, req *tmallchannel.TmallchanneltradeordergetsAPIRequest, session string) (*tmallchannel.TmallchanneltradeordergetsAPIResponse, error) {
	var resp tmallchannel.TmallchanneltradeordergetsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
