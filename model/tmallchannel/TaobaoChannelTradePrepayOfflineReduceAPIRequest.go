package tmallchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoChannelTradePrepayOfflineReduceAPIRequest
渠道分销供应商上传线下流水预存款（减少） API请求
taobao.channel.trade.prepay.offline.reduce

渠道分销供应商上传线下流水预存款（减少） */
type TaobaoChannelTradePrepayOfflineReduceAPIRequest struct {
	model.Params
	// 减少流水
	_offlineReducePrepayParam *TopOfflineReducePrepayDto
}

// New
